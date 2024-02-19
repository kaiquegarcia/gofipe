package selectvehicletype

import "github.com/kaiquegarcia/gofipe/fipe/enums"

type item enums.VehicleType

func (i item) FilterValue() string {
	return i.Label()
}

func (i item) Label() string {
	switch i {
	case item(enums.Car):
		return "Carros"
	case item(enums.Motorbike):
		return "Motos"
	case item(enums.Truck):
		return "Caminhões"
	}

	return "Sair"
}
