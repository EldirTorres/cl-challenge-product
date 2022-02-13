package service

import (
	"cl.falabella.product/src/model"
)

type Service interface {
	saveProduct(product *model.Product)
	updateProduct(product *model.Product)
	removeProduct(id string)
	listProducts() *model.Products
}
