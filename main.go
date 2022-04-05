package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func getClipValue() string {
	cmd := exec.Command("pbpaste")

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Start()
    if err != nil {
		log.Fatal(err)
        os.Exit(1)
	}

	buf := bufio.NewReader(stdout) // Notice that this is not in a loop

	line, _, _ := buf.ReadLine()

	return string(line)
}

func main() {
    p := tea.NewProgram(initialModel())

	// Every Second
	go func() {
		for {
			pause := time.Duration(1 * time.Second)
			time.Sleep(pause)

			p.Send(getClipValue())
		}
	}()

    if err := p.Start(); err != nil {
        fmt.Printf("Alas, there's been an error: %v", err)
        os.Exit(1)
    }
}