package main

import (
	"fmt"
	"net/http"

	"github.com/kaiquegarcia/gofipe/fipe"
	"github.com/kaiquegarcia/gofipe/internal/views/searchvehicle"
	"github.com/kaiquegarcia/gofipe/internal/views/selectbrand"
	"github.com/kaiquegarcia/gofipe/internal/views/selectmodel"
	"github.com/kaiquegarcia/gofipe/internal/views/selectvehicletype"
	"github.com/kaiquegarcia/gofipe/internal/views/selectversion"
)

func main() {
	// clients
	httpClient := http.DefaultClient
	fipeClient := fipe.New(httpClient, fipe.DEFAULT_BASE_URL)

	// navigate
	for {
		vehicleType, err := selectvehicletype.Run()
		if err != nil {
			fmt.Printf("could not select vehicle type: %s\n", err)
			return
		}

		if vehicleType == "" {
			return
		}

		brand, err := selectbrand.Run(fipeClient, vehicleType)
		if err != nil {
			fmt.Printf("could not select vehicle brand: %s\n", err)
			return
		}

		if brand == nil {
			continue
		}

		model, err := selectmodel.Run(fipeClient, vehicleType, brand.Code)
		if err != nil {
			fmt.Printf("could not select vehicle model: %s\n", err)
			return
		}

		if model == nil {
			continue
		}

		version, err := selectversion.Run(fipeClient, vehicleType, brand.Code, model.Code.String())
		if err != nil {
			fmt.Printf("could not select vehicle version: %s\n", err)
			return
		}

		if version == nil {
			continue
		}

		err = searchvehicle.Run(fipeClient, vehicleType, brand.Code, model.Code.String(), version.Code)
		if err != nil {
			fmt.Printf("could not search vehicle: %s\n", err)
			return
		}
	}
}
