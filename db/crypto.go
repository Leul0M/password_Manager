package db

import (
    "bufio"
    "crypto/aes"
    "crypto/cipher"
    "crypto/hmac"
    "crypto/rand"
    "crypto/sha256"
    "encoding/base64"
    "errors"
    "fmt"
    "io"
    "os"
    "path/filepath"

    "golang.org/x/crypto/scrypt"
    "golang.org/x/term"
)

var (
    masterKey []byte
)

const (
    keyFileName  = "vault.key"
    verifyPhrase = "password_manager_v1_verify"
)

// InitCrypto ensures a master key is available by deriving it from a user-provided passphrase.
// On first run, it will create a key file with a random salt and a verification tag.
func InitCrypto() error {
    keyPath := filepath.Join(".", keyFileName)
    if _, err := os.Stat(keyPath); errors.Is(err, os.ErrNotExist) {
        // First run: set a new passphrase and create key file
        pass, err := promptSecret("Create master passphrase: ")
        if err != nil {
            return err
        }
        confirm, err := promptSecret("Confirm master passphrase: ")
        if err != nil {
            return err
        }
        if pass != confirm {
            return errors.New("passphrases do not match")
        }

        salt := make([]byte, 16)
        if _, err := io.ReadFull(rand.Reader, salt); err != nil {
            return err
        }
        key, err := deriveKey([]byte(pass), salt)
        if err != nil {
            return err
        }
        // Store salt + HMAC(verifyPhrase)
        mac := hmac.New(sha256.New, key)
        mac.Write([]byte(verifyPhrase))
        tag := mac.Sum(nil)

        payload := append(salt, tag...)
        if err := os.WriteFile(keyPath, payload, 0600); err != nil {
            return err
        }
        masterKey = key
        return nil
    }

    // Existing key file: prompt for passphrase and verify
    payload, err := os.ReadFile(keyPath)
    if err != nil {
        return err
    }
    if len(payload) < 16+32 {
        return errors.New("invalid key file")
    }
    salt := payload[:16]
    tag := payload[16:48]

    pass, err := promptSecret("Enter master passphrase: ")
    if err != nil {
        return err
    }
    key, err := deriveKey([]byte(pass), salt)
    if err != nil {
        return err
    }
    mac := hmac.New(sha256.New, key)
    mac.Write([]byte(verifyPhrase))
    expected := mac.Sum(nil)
    if !hmac.Equal(tag, expected) {
        return errors.New("invalid master passphrase")
    }
    masterKey = key
    return nil
}

func deriveKey(passphrase, salt []byte) ([]byte, error) {
    // scrypt parameters: N=32768, r=8, p=1 (balanced for CLI)
    return scrypt.Key(passphrase, salt, 32768, 8, 1, 32)
}

func promptSecret(prompt string) (string, error) {
    // Fallback to masked input via term.ReadPassword on stdin
    fmt.Fprint(os.Stdout, prompt)
    fd := int(os.Stdin.Fd())
    if term.IsTerminal(fd) {
        b, err := term.ReadPassword(fd)
        fmt.Fprintln(os.Stdout)
        if err != nil {
            return "", err
        }
        return string(b), nil
    }
    // Non-tty fallback (e.g., piped input) – read a line
    reader := bufio.NewReader(os.Stdin)
    s, err := reader.ReadString('\n')
    if err != nil {
        return "", err
    }
    return s, nil
}

// EncryptString encrypts a plaintext string using AES-GCM and returns base64(nonce||ciphertext)
func EncryptString(plaintext string) (string, error) {
    if masterKey == nil {
        return "", errors.New("crypto not initialized")
    }
    block, err := aes.NewCipher(masterKey)
    if err != nil {
        return "", err
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", err
    }
    ct := gcm.Seal(nil, nonce, []byte(plaintext), nil)
    out := append(nonce, ct...)
    return base64.StdEncoding.EncodeToString(out), nil
}

// DecryptString decrypts base64(nonce||ciphertext) produced by EncryptString
func DecryptString(b64 string) (string, error) {
    if masterKey == nil {
        return "", errors.New("crypto not initialized")
    }
    raw, err := base64.StdEncoding.DecodeString(b64)
    if err != nil {
        return "", err
    }
    block, err := aes.NewCipher(masterKey)
    if err != nil {
        return "", err
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }
    if len(raw) < gcm.NonceSize() {
        return "", errors.New("ciphertext too short")
    }
    nonce := raw[:gcm.NonceSize()]
    ct := raw[gcm.NonceSize():]
    pt, err := gcm.Open(nil, nonce, ct, nil)
    if err != nil {
        return "", err
    }
    return string(pt), nil
}


