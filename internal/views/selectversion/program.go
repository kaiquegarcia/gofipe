package selectversion

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
	modelID string,
) (*entity.Version, error) {
	loading.Start()
	getBrandModelVersionsOutput, err := fipeClient.GetBrandModelVersions(fipe.NewGetBrandModelVersionsInput(
		vehicleType,
		brandID,
		modelID,
	))
	if err != nil {
		return nil, err
	}
	model := newModel(*getBrandModelVersionsOutput.Response)
	loading.Stop()

	program := tea.NewProgram(model, tea.WithAltScreen())
	program.SetWindowTitle("#gofipe > select version")
	if _, err := program.Run(); err != nil {
		return nil, err
	}

	return model.choice.Data(), nil
}
