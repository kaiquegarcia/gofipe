package entity

import "encoding/json"

type Vehicle struct {
	Vehicletype      int         `json:"TipoVeiculo"`
	Price            string      `json:"Valor"`
	Brand            string      `json:"Marca"`
	Model            string      `json:"Modelo"`
	Year             json.Number `json:"AnoModelo"`
	FuelType         string      `json:"Combustivel"`
	FuelAbbreviation string      `json:"SiglaCombustivel"`
	FipeCode         string      `json:"CodigoFipe"`
	ReferenceDate    string      `json:"MesReferencia"`
}
