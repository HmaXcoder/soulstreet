package model

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price float32 `json:"price"`
	Images []string `json:"images"`
	IsAvaliable bool `json:"is_avaliable"`
}