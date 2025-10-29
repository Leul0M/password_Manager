# 🛡️ Password Manager

Welcome to the **Password Manager**! This is a stylish and secure CLI tool built with Go for managing your passwords with ease. 

---

## 🎨 Features
- **Add** new passwords
- **Delete** passwords
- **List** all saved passwords
- **Encryption-at-rest** using AES-GCM with a master passphrase
- **Simple menu navigation**

---

## 🚀 Getting Started

### 1. Clone the repository
```powershell
# Clone the repo
git clone https://github.com/Leul0M/password_Manager.git
cd password_Manager
```

### 2. Build & Run
```powershell
# Build the project
go build -o password_manager.exe

# Run the program
./password_manager.exe
```

On first run, you will be prompted to create a **master passphrase**. This passphrase derives an encryption key using `scrypt`. A local key file `vault.key` (contains only a random salt and verification tag) is created. Keep your passphrase safe—without it, your data cannot be decrypted.

---

## 📋 Usage

When you run the program, you'll see a colorful menu:

- `Add` ➕ : Add a new password
- `Delete` 🗑️ : Remove a password
- `List` 📜 : View all passwords
- `Exit` 🚪 : Quit the program

Just follow the prompts to manage your passwords securely!

### 🔐 Encryption Details
- Your passwords and notes are encrypted before being stored in the SQLite database using **AES-256-GCM**.
- The encryption key is derived from your master passphrase via **scrypt (N=32768, r=8, p=1)** with a 16-byte random salt.
- The key file `vault.key` stores the salt and a verification tag (HMAC) to validate your passphrase on subsequent runs. No plaintext secrets are stored.
- Changing the master passphrase is not yet supported; if needed, export, re-encrypt, and re-import.

---

## 🗄️ Project Structure
```
password_Manager/
│   main.go           # Entry point
│   go.mod, go.sum    # Go modules
│
├── cmd/              # CLI commands
│   ├── add.go        # Add password
│   ├── delete.go     # Delete password
│   ├── list.go       # List passwords
│   ├── menu.go       # Menu logic
│   ├── root.go       # Root command
│   └── style.go      # Colorful styles
│
└── db/
    ├── database.go   # Database logic (encrypts/decrypts fields)
    └── crypto.go     # Key derivation and AES-GCM helpers
```

---

## 🎨 Colors & Style
This app uses colors and emojis to make your experience fun and easy to navigate. The menu and prompts are styled for clarity and visual appeal.

---

## 🛠️ Requirements
- Go 1.18+
- Windows, Mac, or Linux

---

## 🤝 Contributing
Pull requests are welcome! For major changes, please open an issue first to discuss what you would like to change.

---

## 📧 Contact
For questions or feedback, reach out to [Leul0M](https://github.com/Leul0M).

---

## ⚡ License
This project is licensed under the MIT License.

---

> **Stay safe, stay stylish!** ✨🔒
