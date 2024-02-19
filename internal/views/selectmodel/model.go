package selectmodel

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaiquegarcia/gofipe/fipe/entity"
	"github.com/kaiquegarcia/gofipe/internal/views/style"
)

type (
	errMsg error
)

type model struct {
	list   list.Model
	choice item
	err    error
}

func newModel(vehicleModels []entity.Model) *model {
	options := make([]list.Item, len(vehicleModels)+1)
	for i, b := range vehicleModels {
		options[i] = item(b)
	}
	options[len(vehicleModels)] = item(entity.Model{Code: "", Name: "Voltar ao início"})
	list := list.New(
		options,
		itemDelegate{},
		30,
		10,
	)
	list.Title = "De qual modelo?"
	list.SetShowStatusBar(false)
	list.SetFilteringEnabled(true)
	list.Styles.Title = style.TitleStyle
	list.Styles.PaginationStyle = style.PaginationStyle
	list.Styles.HelpStyle = style.HelpStyle

	return &model{
		list: list,
		err:  nil,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			value, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = value
			}

			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return "\n" + m.list.View()
}
