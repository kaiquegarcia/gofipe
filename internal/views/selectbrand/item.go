package selectbrand

import (
	"github.com/kaiquegarcia/gofipe/fipe/entity"
)

type item entity.Brand

func (i item) FilterValue() string {
	return i.Name
}

func (i item) Data() *entity.Brand {
	if i.Code == "" {
		return nil
	}

	return &entity.Brand{
		Code: i.Code,
		Name: i.Name,
	}
}
