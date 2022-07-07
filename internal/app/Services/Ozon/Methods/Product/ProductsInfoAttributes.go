package Product

import (
	"encoding/json"
	"io"
	"main/internal/app/Services/Ozon/Methods"
	"main/internal/app/Services/Ozon/Structures"
	"main/internal/app/Services/Ozon/Structures/Product"
)

// ProductInfoAttributes Информация о товарах/**
func ProductInfoAttributes(body Product.ProductInfoAttributesBody) (*Product.ProductsInfoAttrResult, error) {
	resp, err := Methods.OzonResponse(body, Structures.OzonMethod{Method: "POST", Url: "/v3/products/info/attributes"})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result := Product.ProductsInfoAttrResult{}
	err = json.NewDecoder(io.MultiReader(resp.Body)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}
