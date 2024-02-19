package selectvehicletype

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaiquegarcia/gofipe/fipe/enums"
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

func newModel() *model {
	list := list.New(
		[]list.Item{
			item(enums.Car),
			item(enums.Motorbike),
			item(enums.Truck),
			item(""),
		},
		itemDelegate{},
		30,
		10,
	)
	list.Title = "O que você quer buscar?"
	list.SetShowStatusBar(false)
	list.SetFilteringEnabled(true)
	list.Styles.Title = style.TitleStyle
	list.Styles.PaginationStyle = style.PaginationStyle
	list.Styles.HelpStyle = style.HelpStyle

	return &model{
		list:   list,
		choice: "",
		err:    nil,
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
