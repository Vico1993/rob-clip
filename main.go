package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)


func main() {
    p := tea.NewProgram(initialModel())

	// Every Second
	go func() {
		for {
			pause := time.Duration(1 * time.Second)
			time.Sleep(pause)

			p.Send(Copyed{
				word: GetValue(),
				date: time.Now(),
			})
		}
	}()

    if err := p.Start(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}