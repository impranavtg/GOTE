package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	msg                   string
	newNoteInput          textinput.Model
	noteTextarea          textarea.Model
	isNewNoteInputVisible bool
	currentNote           *os.File
	noteList              list.Model
	isNoteListVisible     bool
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.noteList.SetSize(msg.Width-h, msg.Height-v-5)

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+n":
			m.isNewNoteInputVisible = true
			return m, nil
		case "ctrl+l":
			gotes := listGotes()
			m.noteList.SetItems(gotes)
			m.isNoteListVisible = true
			return m, nil
		case "ctrl+s":
			if m.currentNote == nil {
				break
			}

			// write the value of the noteTextarea into the currentNote and close it
			if err := m.currentNote.Truncate(0); err != nil {
				fmt.Println("Cannot save your gote idea 😶")
				return m, nil
			}

			if _, err := m.currentNote.Seek(0, 0); err != nil {
				fmt.Println("Cannot save your gote idea 😶")
				return m, nil
			}

			if _, err := m.currentNote.WriteString(m.noteTextarea.Value()); err != nil {
				fmt.Println("Cannot save your gote idea 😶")
				return m, nil
			}

			if err := m.currentNote.Close(); err != nil {
				fmt.Println("Cannot close the file")
			}

			m.currentNote = nil
			m.noteTextarea.Reset()

			return m, nil

		case "esc":
			if m.isNewNoteInputVisible {
				m.isNewNoteInputVisible = false
			}
			if m.currentNote != nil {
				m.currentNote = nil
				m.noteTextarea.Reset()
			}
			if m.isNoteListVisible {
				if m.noteList.FilterState() == list.Filtering {
					break
				}
				m.isNoteListVisible = false
			}
		case "enter":
			if m.currentNote != nil {
				break
			}
			if m.isNoteListVisible {
				selectedGote, ok := m.noteList.SelectedItem().(gote)
				if ok {
					filePath := filepath.Join(notesDir, selectedGote.title)
					content, err := os.ReadFile(filePath)
					if err != nil {
						log.Fatal("Error reading the Gote 😶")
						return m, nil
					}
					m.noteTextarea.SetValue(string(content))
					f, err := os.OpenFile(filePath, os.O_RDWR, 0644)
					if err != nil {
						log.Fatal("Error reading the Gote 😶")
						return m, nil
					}
					m.currentNote = f

					m.isNoteListVisible = false
				}
				return m, nil
			}
			fileName := m.newNoteInput.Value()
			if fileName != "" {
				filePath := filepath.Join(notesDir, fileName+".md")
				if _, err := os.Stat(filePath); err == nil {
					fmt.Println("File already exists")
					return m, nil
				}
				f, err := os.Create(filePath)
				if err != nil {
					fmt.Println("Error creating file:", err)
					return m, nil
				}

				m.currentNote = f
				m.isNewNoteInputVisible = false
				m.newNoteInput.Reset()

			}
			return m, nil

		}

	}

	if m.isNewNoteInputVisible {
		m.newNoteInput, cmd = m.newNoteInput.Update(msg)
	}
	if m.currentNote != nil {
		m.noteTextarea, cmd = m.noteTextarea.Update(msg)
	}

	if m.isNoteListVisible {
		m.noteList, cmd = m.noteList.Update(msg)
	}

	return m, cmd
}

func (m model) View() string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("220")).
		Padding(0, 1)

	var helpMsgStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("244")).
		Padding(0, 1)

	welcomeMsg := style.Render("Welcome to Gote 📝")
	view := ""
	if m.isNewNoteInputVisible {
		view = m.newNoteInput.View()
	}
	if m.currentNote != nil {
		view = m.noteTextarea.View()
	}
	if m.isNoteListVisible {
		view = m.noteList.View()
	}
	helpMsg := helpMsgStyle.Render("Ctrl+N: New Note | Ctrl+L: List Notes | Esc: Back | Ctrl+S: Save Note | Ctrl+C: Quit")
	return fmt.Sprintf("\n%s\n\n%s\n\n%s\n\n", welcomeMsg, view, helpMsg)
}

func initialModel() model {
	// create a folder at the home directory of the user with the name .gote
	err := os.MkdirAll(notesDir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
	}

	ti := textinput.New()
	ti.Placeholder = "What's your gote?"
	ti.Focus()
	ti.CharLimit = 156
	ti.Cursor.Style = cursorStyle
	ti.PromptStyle = cursorStyle
	ti.TextStyle = cursorStyle

	ta := textarea.New()
	ta.Placeholder = "Write your gote idea"
	ta.Focus()
	ta.ShowLineNumbers = false

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	items := listGotes()

	itemList := list.New(items, list.NewDefaultDelegate(), 0, 0)
	itemList.Title = "> Gote List 🗒️"
	itemList.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("220")).
		Background(lipgloss.Color("16")).
		Padding(0, 1)

	return model{
		msg:                   "Welcome to Gote App!!",
		newNoteInput:          ti,
		isNewNoteInputVisible: false,
		noteTextarea:          ta,
		noteList:              itemList,
		isNoteListVisible:     false,
	}

}
