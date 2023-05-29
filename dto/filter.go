package dto

type Filter struct {
	Name       string `query:"name"`
	Barcode    string `query:"barcode"`
	Price      uint   `query:"price"`
	CategoryID uint   `query:"category_id"`
}
