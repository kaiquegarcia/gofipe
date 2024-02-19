package entity

import "encoding/json"

type Model struct {
	Code json.Number `json:"codigo"`
	Name string      `json:"nome"`
}

type GetModelsResponse struct {
	Models   []Model   `json:"modelos"`
	Versions []Version `json:"anos"`
}
