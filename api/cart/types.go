package cart

import "shopping_go/domain/cart"

type ItemCartRequest struct {
	SKU   string `json:"sku"`
	Count int    `json:"count"`
}

type CreateCategoryResponse struct {
	Message string `json:"message"`
}

type ItemCartResponse struct {
	items []cart.Item `json:"items"`
}
