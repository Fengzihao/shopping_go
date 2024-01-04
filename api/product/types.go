package product

import "shopping_go/domain/product"

type CreateProductRequest struct {
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Price      float32 `json:"price"`
	Count      int     `json:"count"`
	CategoryID uint    `json:"categoryID"`
}

type CreateProductResponse struct {
	Message string `json:"message"`
}

type DeleteProductRequest struct {
	SKU string `json:"sku"`
}

// 更新product的request
type UpdateProductRequest struct {
	SKU        string  `json:"sku"`
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Price      float32 `json:"price"`
	Count      int     `json:"count"`
	CategoryID uint    `json:"categoryID"`
}

// 实体转换
func (receiver *UpdateProductRequest) ToProduct() *product.Product {
	return product.NewProduct(receiver.Name, receiver.Desc, receiver.Count, receiver.Price, receiver.CategoryID)
}
