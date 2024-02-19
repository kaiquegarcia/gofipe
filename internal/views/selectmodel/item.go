package selectmodel

import (
	"github.com/kaiquegarcia/gofipe/fipe/entity"
)

type item entity.Model

func (i item) FilterValue() string {
	return i.Name
}

func (i item) Data() *entity.Model {
	if i.Code == "" {
		return nil
	}

	return &entity.Model{
		Code: i.Code,
		Name: i.Name,
	}
}
