package name

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"music-project-scaffolder-bubbletea/cmd/styles"
)

type ProjectName struct {
	ti   textinput.Model
	name string
}

var textStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("170"))
var promptStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))

func InitProjectName() ProjectName {
	ti := textinput.New()
	ti.Placeholder = "My awesome project"
	ti.CharLimit = 156
	ti.Width = 30
	ti.TextStyle = textStyle
	ti.CursorStyle = lipgloss.NewStyle().Blink(true)
	ti.PromptStyle = promptStyle

	return ProjectName{
		ti:   ti,
		name: "",
	}
}

func (n ProjectName) Init() tea.Cmd {
	//return textinput.Blink
	return nil
}

func (n ProjectName) Update(msg tea.Msg) (mod tea.Model, cmd tea.Cmd) {
	n.ti.Focus()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", "tab", "esc":
			n.ti.Blur()
			return n, func() tea.Msg {
				return Finished("finished")
			}
		}
	}
	n.ti, cmd = n.ti.Update(msg)
	return n, cmd
}

func (n ProjectName) View() string {
	return fmt.Sprintf(
		styles.TitleStyle.Render("What is the name of your project?")+
			"\n\n"+
			"%s\n\n",
		n.ti.View(),
	)
}
