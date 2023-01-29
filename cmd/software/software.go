package software

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"music-project-scaffolder-bubbletea/cmd/styles"
)

type ProjectSoftware struct {
	choices  []Software
	cursor   int
	selected map[int]Software
}

func InitiateSoftware() ProjectSoftware {
	return ProjectSoftware{
		choices:  []Software{FlStudio, Ableton, ProTools},
		selected: make(map[int]Software),
	}
}

func (s ProjectSoftware) Init() tea.Cmd {
	return nil
}

func (s ProjectSoftware) Update(msg tea.Msg) (mod tea.Model, cmd tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if s.cursor > 0 {
				s.cursor--
			}
		case "down", "j":
			if s.cursor < len(s.choices)-1 {
				s.cursor++
			}
		case " ":
			_, ok := s.selected[s.cursor]
			if ok {
				delete(s.selected, s.cursor)
			} else {
				s.selected[s.cursor] = s.choices[s.cursor]
			}
		case "enter":
			return s, func() tea.Msg {
				return Finished("finished")
			}

		}

	}
	return s, cmd
}

func (s ProjectSoftware) View() string {
	str := styles.TitleStyle.Render("Which software do you want to use?") + "\n\n"

	for i, choice := range s.choices {
		cursor := " "
		style := inactiveSelectionStyle
		if s.cursor == i {
			cursor = ">"
			style = activeSelectionStyle
		}
		checked := " "
		if _, ok := s.selected[i]; ok {
			checked = "x"
		}
		str += style.Render(fmt.Sprintf("%s [%s] %s", cursor, checked, choice)) + "\n"
	}

	str += "\n" + helpText.Render("Use ") + keysStyle.Render("\"space\"") +
		helpText.Render(" to select and ") + keysStyle.Render("\"enter\"") +
		helpText.Render(" to finish")

	//str += "\nPress q to quit.\n"
	return str
}

var (
	helpText               = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	keysStyle              = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	activeSelectionStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("170")).PaddingLeft(2)
	inactiveSelectionStyle = lipgloss.NewStyle().PaddingLeft(2)
)
