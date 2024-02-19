package searchvehicle

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaiquegarcia/gofipe/fipe/entity"
)

type (
	errMsg error
)

type model struct {
	vehicle entity.Vehicle
	err     error
}

func newModel(vehicle entity.Vehicle) *model {
	return &model{
		vehicle: vehicle,
		err:     nil,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc, tea.KeyEnter:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	return fmt.Sprintf(
		`
 Veículo encontrado:
  Código FIPE: %s
  Marca: %s
  Modelo: %s  Ano: %s
  Valor: %s
  Tipo de Combustível: %s`,
		m.vehicle.FipeCode,
		m.vehicle.Brand,
		m.vehicle.Model, m.vehicle.Year.String(),
		m.vehicle.Price,
		m.vehicle.FuelType,
	)
}
