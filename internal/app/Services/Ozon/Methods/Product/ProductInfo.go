package Product

import (
	"encoding/json"
	"io"
	"main/internal/app/Services/Ozon/Methods"
	"main/internal/app/Services/Ozon/Structures"
	"main/internal/app/Services/Ozon/Structures/Product"
)

// ProductInfo Информация о товарах/**
func ProductInfo(body Product.ProductInfoBody) (*Product.ProductInfoResult, error) {
	resp, err := Methods.OzonResponse(body, Structures.OzonMethod{Method: "POST", Url: "/v2/product/info"})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result := Product.ProductInfoResult{}
	err = json.NewDecoder(io.MultiReader(resp.Body)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}
