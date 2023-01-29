package _type

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"io"
)

type ProjectType struct {
	options  list.Model
	choice   Type
	quitting bool
}

func InitProjectType() ProjectType {
	items := []list.Item{
		Beat,
		Song,
		Remix,
		Voiceover,
		Edit,
	}
	const defaultWidth = 40
	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "What kind of project is it?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	return ProjectType{
		options:  l,
		quitting: false,
	}
}

func (t ProjectType) Init() tea.Cmd {
	return nil
}

func (t ProjectType) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		t.options.SetWidth(msg.Width)
		return t, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "enter":
			i, ok := t.options.SelectedItem().(Type)
			if ok {
				t.choice = i
				t.quitting = true
			}
			return t, func() tea.Msg {
				return Finished("finished")
			}
		}
	}

	var cmd tea.Cmd
	t.options, cmd = t.options.Update(msg)
	return t, cmd
}

func (t ProjectType) View() string {
	if t.choice != 0 {
		return "FinishedText" + t.choice.String()
	}
	return "\n" + t.options.View()
}

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(0).Bold(true)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Type)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.String())

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprint(w, fn(str))
}
