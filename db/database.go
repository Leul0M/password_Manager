package db

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Credential struct {
	ID       int
	Email    string
	Password string
	Note     string
}

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./vault.db")
	if err != nil {
		return err
	}

	query := `
	CREATE TABLE IF NOT EXISTS credentials (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT,
		password TEXT,
		note TEXT
		)`
		_, err = DB.Exec(query)
		return err
}

func AddCredential(email, password, note string) error {
    encPassword, err := EncryptString(password)
    if err != nil {
        return err
    }
    encNote, err := EncryptString(note)
    if err != nil {
        return err
    }
    _, err = DB.Exec("INSERT INTO credentials (email, password, note) VALUES (?, ?, ?)", email, encPassword, encNote)
	return err
}

func GetAllCredentials() ([]Credential, error) {
    rows, err := DB.Query("SELECT id, email, password, note FROM credentials")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var creds []Credential
	for rows.Next() {
		var c Credential
        var encPass, encNote string
        err := rows.Scan(&c.ID, &c.Email, &encPass, &encNote)
		if err != nil {
			return nil, err
		}
        if dec, derr := DecryptString(encPass); derr == nil {
            c.Password = dec
        } else {
            c.Password = "<decrypt error>"
        }
        if decN, derr := DecryptString(encNote); derr == nil {
            c.Note = decN
        } else {
            c.Note = "<decrypt error>"
        }
		creds = append(creds, c)
	}
	return creds, nil
}


func ListCredentials() (string, error) {
    rows, err := DB.Query("SELECT id, email, password, note FROM credentials")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var result strings.Builder
	for rows.Next() {
		var id int
        var email, encPass, encNote string
        err = rows.Scan(&id, &email, &encPass, &encNote)
		if err != nil {
			return "", err
		}
        decPass, derr1 := DecryptString(encPass)
        if derr1 != nil { decPass = "<decrypt error>" }
        decNote, derr2 := DecryptString(encNote)
        if derr2 != nil { decNote = "<decrypt error>" }
        result.WriteString(fmt.Sprintf("[%d] Email: %s | Password: %s | Note: %s\n", id, email, decPass, decNote))
	}

	return result.String(), nil
}

func DeleteCredential(id int) error {
	_, err := DB.Exec("DELETE FROM credentials WHERE id = ?", id)
	return err
}
