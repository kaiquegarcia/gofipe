package fipe

import (
	"fmt"
	"net/http"

	"github.com/kaiquegarcia/gofipe/fipe/entity"
	"github.com/kaiquegarcia/gofipe/fipe/enums"
)

// GetBrandModelVersionsInput contains all the possible inputs for the versions' search query.
//
// Note: avoid initializing it on your own. Use NewGetBrandModelVersionsInput() instead.
type GetBrandModelVersionsInput struct {
	vehicleType enums.VehicleType
	brandID     string
	modelID     string
}

// NewGetBrandModelVersionsInput returns an initilized input struct to compose the versions' search query.
func NewGetBrandModelVersionsInput(
	vehicleType enums.VehicleType,
	brandID string,
	modelID string,
) GetBrandModelVersionsInput {
	return GetBrandModelVersionsInput{
		vehicleType: vehicleType,
		brandID:     brandID,
		modelID:     modelID,
	}
}

type GetBrandModelVersionsOutput struct {
	Audit    *entity.Audit
	Response *entity.GetVersionsResponse
}

func (f *client) GetBrandModelVersions(i GetBrandModelVersionsInput) (*GetBrandModelVersionsOutput, error) {
	var response entity.GetVersionsResponse
	audit, err := f.doRequest(
		http.MethodGet,
		fmt.Sprintf("%s/marcas/%s/modelos/%s/anos", i.vehicleType, i.brandID, i.modelID),
		[]byte{},
		&response,
	)

	return &GetBrandModelVersionsOutput{
		Audit:    &audit,
		Response: &response,
	}, err
}
