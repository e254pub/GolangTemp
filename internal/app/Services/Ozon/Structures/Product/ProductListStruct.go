package Product

import (
	"main/internal/app/Services/Ozon/Structures"
)

type ProductListBody struct {
	Filter Structures.Filter `json:"filter"`
	LastId string            `json:"last_id"`
	Limit  int               `json:"limit"`
}

type ProductsListResult struct {
	Result ProductListResult `json:"result"`
}

type ProductListResult struct {
	Items []struct {
		OfferId   string `json:"offer_id"`
		ProductId int    `json:"product_id"`
	} `json:"items"`
	LastId string `json:"last_id"`
	Total  int    `json:"total"`
}
