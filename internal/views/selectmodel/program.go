package selectmodel

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kaiquegarcia/gofipe/fipe"
	"github.com/kaiquegarcia/gofipe/fipe/entity"
	"github.com/kaiquegarcia/gofipe/fipe/enums"
	"github.com/kaiquegarcia/gofipe/internal/views/loading"
)

func Run(
	fipeClient fipe.Client,
	vehicleType enums.VehicleType,
	brandID string,
) (*entity.Model, error) {
	loading.Start()
	getBrandModelsOutput, err := fipeClient.GetBrandModels(fipe.NewGetBrandModelsInput(
		vehicleType,
		brandID,
	))
	if err != nil {
		return nil, err
	}
	model := newModel(getBrandModelsOutput.Response.Models)
	loading.Stop()

	program := tea.NewProgram(model, tea.WithAltScreen())
	program.SetWindowTitle("#gofipe > select model")
	if _, err := program.Run(); err != nil {
		return nil, err
	}

	return model.choice.Data(), nil
}
