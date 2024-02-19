package fipe

import (
	"fmt"
	"net/http"

	"github.com/kaiquegarcia/gofipe/fipe/entity"
	"github.com/kaiquegarcia/gofipe/fipe/enums"
)

// GetBrandsInput contains all the possible inputs for the brands' search query.
//
// Note: avoid initializing it on your own. Use NewGetBrandsInput() instead.
type GetBrandsInput struct {
	vehicleType enums.VehicleType
}

// NewGetBrandsInput returns an initilized input struct to compose the brands' search query.
func NewGetBrandsInput(vehicleType enums.VehicleType) GetBrandsInput {
	return GetBrandsInput{
		vehicleType: vehicleType,
	}
}

type GetBrandsOutput struct {
	Audit    *entity.Audit
	Response *entity.GetBrandsResponse
}

func (f *client) GetBrands(i GetBrandsInput) (*GetBrandsOutput, error) {
	var response entity.GetBrandsResponse
	audit, err := f.doRequest(
		http.MethodGet,
		fmt.Sprintf("%s/marcas", i.vehicleType),
		[]byte{},
		&response,
	)

	return &GetBrandsOutput{
		Audit:    &audit,
		Response: &response,
	}, err
}
