package selectvehicletype

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaiquegarcia/gofipe/fipe/enums"
)

func Run() (enums.VehicleType, error) {
	model := newModel()
	program := tea.NewProgram(model, tea.WithAltScreen())
	program.SetWindowTitle("#gofipe > select vehicle type")
	if _, err := program.Run(); err != nil {
		return "", err
	}

	return enums.VehicleType(model.choice), nil
}
