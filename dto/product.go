package dto

type CProduct struct {
	Name       string `json:"name"`
	Barcode    string `json:"barcode"`
	Price      uint   `json:"price"`
	CategoryID uint   `json:"category_id"`
}

type UProduct struct {
	Name  string `json:"name"`
	Price uint   `json:"price"`
}
