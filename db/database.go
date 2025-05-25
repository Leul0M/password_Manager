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
	_, err := DB.Exec("INSERT INTO credentials (email, password, note) VALUES (?, ?, ?)", email, password, note)
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
		err := rows.Scan(&c.ID, &c.Email, &c.Password, &c.Note)
		if err != nil {
			return nil, err
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
		var email, password, note string
		err = rows.Scan(&id, &email, &password, &note)
		if err != nil {
			return "", err
		}
		result.WriteString(fmt.Sprintf("[%d] Email: %s | Password: %s | Note: %s\n", id, email, password, note))
	}

	return result.String(), nil
}

func DeleteCredential(id int) error {
	_, err := DB.Exec("DELETE FROM credentials WHERE id = ?", id)
	return err
}
