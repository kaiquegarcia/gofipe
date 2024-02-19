package selectbrand

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
) (*entity.Brand, error) {
	loading.Start()
	getBrandsOutput, err := fipeClient.GetBrands(fipe.NewGetBrandsInput(vehicleType))
	if err != nil {
		return nil, err
	}
	model := newModel(*getBrandsOutput.Response)
	loading.Stop()

	program := tea.NewProgram(model, tea.WithAltScreen())
	program.SetWindowTitle("#gofipe > select brand")
	if _, err := program.Run(); err != nil {
		return nil, err
	}

	return model.choice.Data(), nil
}
