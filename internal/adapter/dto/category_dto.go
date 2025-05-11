package dto

type DtoCreateCategory struct {
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	UUID         string `json:"uuid"`
}
