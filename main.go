package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"passwordmanager/cmd"
	"passwordmanager/db"
)

func main() {
	err := db.InitDB()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}

	// Start Bubbletea program with main menu model
	p := tea.NewProgram(cmd.NewMainMenuModel())
	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
