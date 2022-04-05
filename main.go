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

type model struct {
    previousCopy  []string
    cursor   int
    selected map[int]struct{}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

        // Is it a key press?
        case tea.KeyMsg:
            // Cool, what was the actual key pressed?
            switch msg.String() {

                // These keys should exit the program.
                case "ctrl+c", "q":
                    return m, tea.Quit

                // The "up" and "k" keys move the cursor up
                case "up", "k":
                    if m.cursor > 0 {
                        m.cursor--
                    }

                // The "down" and "j" keys move the cursor down
                case "down", "j":
                    if m.cursor < len(m.previousCopy)-1 {
                        m.cursor++
                    }

                // The "enter" key and the spacebar (a literal space) toggle
                // the selected state for the item that the cursor is pointing at.
                case "enter", " ":
                    _, ok := m.selected[m.cursor]
                    if ok {
                        delete(m.selected, m.cursor)
                    } else {
                        m.selected[m.cursor] = struct{}{}
                    }
            }

        case string:
            if (len(m.previousCopy) == 0) {
                m.previousCopy = append(m.previousCopy, msg)
            } else {
                latestCopy := m.previousCopy[len(m.previousCopy) -1]

                if (latestCopy != msg) {
                    m.previousCopy = append(m.previousCopy, msg)
                }
            }
	}

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return m, nil
}

func (m model) View() string {
    // The header
    s := "What did I copy before?\n\n"

    // Iterate over our choices
    for i, choice := range m.previousCopy {

        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
        }

        // Is this choice selected?
        checked := " " // not selected
        if _, ok := m.selected[i]; ok {
            checked = "x" // selected!
        }

        // Render the row
        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    return s
}

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