package cmd

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"music-project-scaffolder-bubbletea/cmd/name"
	software2 "music-project-scaffolder-bubbletea/cmd/software"
	_type "music-project-scaffolder-bubbletea/cmd/type"
	"os"
)

type SessionState int

const (
	projectName SessionState = iota
	projectType
	software
	create
)

type MainModel struct {
	state       SessionState
	name        tea.Model
	projectType tea.Model
	software    tea.Model
}

type projectConfig struct {
	name        string
	projectType string
	software    []string
}

func initialModel() MainModel {
	return MainModel{
		name:        name.InitProjectName(),
		projectType: _type.InitProjectType(),
		software:    software2.InitiateSoftware(),
	}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (mod tea.Model, cmd tea.Cmd) {
	switch amsg := msg.(type) {
	case tea.KeyMsg:
		switch amsg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	case name.Finished:
		m.state = projectType
	case _type.Finished:
		m.state = software
	case software2.Finished:
		m.state = create

	}
	switch m.state {
	case projectName:
		newName, newCmd := m.name.Update(msg)
		nameModel, ok := newName.(name.ProjectName)
		if !ok {
			panic("Could not assert nameModel")
		}
		m.name = nameModel
		cmd = newCmd
	case projectType:
		newType, newCmd := m.projectType.Update(msg)
		typeModel, ok := newType.(_type.ProjectType)
		if !ok {
			panic("could not assert typeModel")
		}
		m.projectType = typeModel
		cmd = newCmd
	case software:
		newSoftware, newCmd := m.software.Update(msg)
		softwareModel, ok := newSoftware.(software2.ProjectSoftware)
		if !ok {
			panic("could not assert Softwaremodel")
		}
		m.software = softwareModel
		cmd = newCmd
	case create:

	}

	return m, cmd
}

func (m MainModel) View() string {
	switch m.state {
	case projectName:
		return m.name.View()
	case projectType:
		return m.projectType.View()
	case software:
		return m.software.View()
	case create:
		return "finished"
	}
	return "problem"
}

func Start() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Problem: %v", err)
		os.Exit(1)
	}
}
