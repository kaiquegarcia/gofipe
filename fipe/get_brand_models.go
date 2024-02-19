package fipe

import (
	"fmt"
	"net/http"

	"github.com/kaiquegarcia/gofipe/fipe/entity"
	"github.com/kaiquegarcia/gofipe/fipe/enums"
)

// GetBrandModelsInput contains all the possible inputs for the models' search query.
//
// Note: avoid initializing it on your own. Use NewGetBrandModelsInput() instead.
type GetBrandModelsInput struct {
	vehicleType enums.VehicleType
	brandID     string
}

// NewGetBrandModelsInput returns an initilized input struct to compose the models' search query.
func NewGetBrandModelsInput(vehicleType enums.VehicleType, brandID string) GetBrandModelsInput {
	return GetBrandModelsInput{
		vehicleType: vehicleType,
		brandID:     brandID,
	}
}

type GetBrandModelsOutput struct {
	Audit    *entity.Audit
	Response *entity.GetModelsResponse
}

func (f *client) GetBrandModels(i GetBrandModelsInput) (*GetBrandModelsOutput, error) {
	var response entity.GetModelsResponse
	audit, err := f.doRequest(
		http.MethodGet,
		fmt.Sprintf("%s/marcas/%s/modelos", i.vehicleType, i.brandID),
		[]byte{},
		&response,
	)

	return &GetBrandModelsOutput{
		Audit:    &audit,
		Response: &response,
	}, err
}
