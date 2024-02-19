package entity

type Version struct {
	Code string `json:"codigo"`
	Name string `json:"nome"`
}

type GetVersionsResponse []Version
