package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Copyed struct {
    word string
    date time.Time
}

type clipboardHistory struct {
    list  []Copyed
    cursor   int
    selected map[int]struct{}
}

func (c clipboardHistory) findWordInList(word string) bool {
    for i := 0; i < len(c.list); i++ {
        if c.list[i].word == word {
            return false
        }
    }

    return true
}

func (c clipboardHistory) navigationKey(key string) (tea.Model, tea.Cmd) {
	// Cool, what was the actual key pressed?
	switch key {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return c, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if c.cursor > 0 {
				c.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if c.cursor < len(c.list)-1 {
				c.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := c.selected[c.cursor]
			if ok {
				delete(c.selected, c.cursor)
			} else {
				c.selected[c.cursor] = struct{}{}
			}
	}

	return c, nil
}

func (c clipboardHistory) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
        // Is it a key press?
        case tea.KeyMsg:
            return c.navigationKey(msg.String())

        case Copyed:
            if (c.findWordInList(msg.word)) {
                c.list = append(c.list, msg)
            }
	}

    // Return the updated model to the Bubble Tea runtime for processing.
    // Note that we're not returning a command.
    return c, nil
}

func (c clipboardHistory) View() string {
    // The header
    s := "The is your clipboard history? 📝\n\n"

    // Iterate over our choices
    for i, choice := range c.list {

        // Is the cursor pointing at this choice?
        cursor := " " // no cursor
        if c.cursor == i {
            cursor = ">" // cursor!
        }

        // Is this choice selected?
        checked := " " // not selected
        if _, ok := c.selected[i]; ok {
            checked = "x" // selected!
        }

        // Render the row
        s += fmt.Sprintf(
            lipgloss.NewStyle().Bold(true).Render("%s [%s] %s"),
            cursor,checked,choice.word,
        )
        s += fmt.Sprintln(
            lipgloss.NewStyle().Italic(true).Render(" -"),
            choice.date.Format("2006-01-02 15:04:05"),
        )
    }

    // The footer
    s += "\nPress q to quit.\n"

    // Send the UI for rendering
    return s
}

func (c clipboardHistory) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}