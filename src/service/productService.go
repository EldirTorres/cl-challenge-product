package service

import (
	"cl.challenge.product/src/model"
)

type Service interface {
	saveProduct(product *model.Product)
	updateProduct(product *model.Product)
	removeProduct(id string)
	listProducts() *model.Products
}
