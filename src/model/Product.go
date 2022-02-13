package model

type Product struct {
	Sku            string   `json:"id"`
	Name           string   `json:"name"`
	Brand          string   `json:"brand"`
	Size           int      `json:"size"`
	Price          float32  `json:"price"`
	PrincipalImage string   `json:"productImage"`
	OtherImages    []string `json:"productImages"`
}

type Products []Product
