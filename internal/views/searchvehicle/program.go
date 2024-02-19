package searchvehicle

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaiquegarcia/gofipe/fipe"
	"github.com/kaiquegarcia/gofipe/fipe/enums"
	"github.com/kaiquegarcia/gofipe/internal/views/loading"
)

func Run(
	fipeClient fipe.Client,
	vehicleType enums.VehicleType,
	brandID string,
	modelID string,
	versionID string,
) error {
	loading.Start()
	getVehicleOutput, err := fipeClient.GetVehicle(fipe.NewGetVehicleInput(
		vehicleType,
		brandID,
		modelID,
		versionID,
	))
	if err != nil {
		return err
	}
	model := newModel(*getVehicleOutput.Response)
	loading.Stop()

	program := tea.NewProgram(model, tea.WithAltScreen())
	program.SetWindowTitle("#gofipe > result")
	_, err = program.Run()
	return err
}
