package fipe

import (
	"fmt"
	"net/http"

	"github.com/kaiquegarcia/gofipe/fipe/entity"
	"github.com/kaiquegarcia/gofipe/fipe/enums"
)

// GetVehicleInput contains all the possible inputs for the vehicles' search query.
//
// Note: avoid initializing it on your own. Use NewGetVehicleInput() instead.
type GetVehicleInput struct {
	vehicleType enums.VehicleType
	brandID     string
	modelID     string
	versionID   string
}

// NewGetVehicleInput returns an initilized input struct to compose the vehicles' search query.
func NewGetVehicleInput(
	vehicleType enums.VehicleType,
	brandID string,
	modelID string,
	versionID string,
) GetVehicleInput {
	return GetVehicleInput{
		vehicleType: vehicleType,
		brandID:     brandID,
		modelID:     modelID,
		versionID:   versionID,
	}
}

type GetVehicleOutput struct {
	Audit    *entity.Audit
	Response *entity.Vehicle
}

func (f *client) GetVehicle(i GetVehicleInput) (*GetVehicleOutput, error) {
	var response entity.Vehicle
	audit, err := f.doRequest(
		http.MethodGet,
		fmt.Sprintf("%s/marcas/%s/modelos/%s/anos/%s", i.vehicleType, i.brandID, i.modelID, i.versionID),
		[]byte{},
		&response,
	)

	return &GetVehicleOutput{
		Audit:    &audit,
		Response: &response,
	}, err
}
