package fipe

import (
	"net/http"
)

// Client is a service responsible to retrieve brazilian's car information present on FIPE table.
//
// The current implementation uses the free API https://deividfortuna.github.io/fipe/ to access the informations freely.
// If someone ends this API, this implementation won't work anymore.
//
// The recommended baseURL is "https://parallelum.com.br/fipe/api/v1", stored in a public constant "DEFAULT_BASE_URL"
type Client interface {
	// GetBrands will look up for brands to compose the first screen: setbrand.
	// This is the first required step to get a vehicle's data, also required for the next step (GetBrandModels)
	GetBrands(GetBrandsInput) (*GetBrandsOutput, error)
	// GetBrandModels will look up for models related to a brandID and compose the second screen: setmodel.
	// This is the second required step to get a vehicle's data, also required for the next step (GetBrandModelVersions).
	GetBrandModels(GetBrandModelsInput) (*GetBrandModelsOutput, error)
	// GetBrandModelVersions will look up for versions related to a modelID from a brandID and compose the third screen: setversion.
	// This is the third required step to get a vehicle's data.
	GetBrandModelVersions(GetBrandModelVersionsInput) (*GetBrandModelVersionsOutput, error)
	// GetVehicle will retrieve a vehicle's information from FIPE table.
	GetVehicle(GetVehicleInput) (*GetVehicleOutput, error)
}

// New returns our private implementation of Client interface
func New(
	httpClient *http.Client,
	baseURL string,
) Client {
	return &client{
		httpClient: httpClient,
		baseURL:    baseURL,
	}
}

const DEFAULT_BASE_URL = "https://parallelum.com.br/fipe/api/v1"

type client struct {
	httpClient *http.Client
	baseURL    string
}
