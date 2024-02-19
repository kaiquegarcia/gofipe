package selectversion

import (
	"github.com/kaiquegarcia/gofipe/fipe/entity"
)

type item entity.Version

func (i item) FilterValue() string {
	return i.Name
}

func (i item) Data() *entity.Version {
	if i.Code == "" {
		return nil
	}

	return &entity.Version{
		Code: i.Code,
		Name: i.Name,
	}
}
