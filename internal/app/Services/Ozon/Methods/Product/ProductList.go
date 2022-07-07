package Product

import (
	"encoding/json"
	"io"
	"main/internal/app/Services/Ozon/Methods"
	"main/internal/app/Services/Ozon/Structures"
	"main/internal/app/Services/Ozon/Structures/Product"
)

// ProductList Список товаров/**
func ProductList(body Product.ProductListBody) (*Product.ProductsListResult, error) {
	resp, err := Methods.OzonResponse(body, Structures.OzonMethod{Method: "POST", Url: "/v2/product/list"})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result := Product.ProductsListResult{}
	err = json.NewDecoder(io.MultiReader(resp.Body)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}
