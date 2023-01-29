package main

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"music-project-scaffolder-bubbletea/cmd/enums"
	"os"
)

type model struct {
	projectName textinput.Model
	choices     []enums.ProjectType
	cursor      int
	selected    map[int]struct{}
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "My awesome project"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 30

	return model{
		projectName: ti,
		choices:     []enums.ProjectType{enums.Beat, enums.Song, enums.Remix, enums.Voiceover, enums.Edit},
		selected:    make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (mod tea.Model, cmd tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			if !m.projectName.Focused() {
				tea.Quit()
			}
		case "tab":
			if m.projectName.Focused() {
				m.projectName.Blur()
			} else {
				tea.Quit()
			}

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			m.selected = make(map[int]struct{})
			m.selected[m.cursor] = struct{}{}

		}

	}
	m.projectName, cmd = m.projectName.Update(msg)
	return m, cmd
}

func (m model) View() string {

	return fmt.Sprintf(
		"What is the name of your project?\n\n"+
			"%s\n\n"+
			"%s\n",
		m.projectName.View(),
		"ctrl-c to quit",
	)

	//for i, choice := range m.choices {
	//	cursor := " "
	//	if m.cursor == i {
	//		cursor = ">"
	//	}
	//	checked := " "
	//	if _, ok := m.selected[i]; ok {
	//		checked = "x"
	//	}
	//	s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice.String())
	//}
	//
	//s += "\nPress q to quit.\n"
	//return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Problem: %v", err)
		os.Exit(1)
	}
}
