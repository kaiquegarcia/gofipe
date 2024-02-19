package entity

type Brand struct {
	Code string `json:"codigo"`
	Name string `json:"nome"`
}

type GetBrandsResponse []Brand
