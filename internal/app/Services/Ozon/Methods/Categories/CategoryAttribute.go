package Categories

import (
	"encoding/json"
	"io"
	"main/internal/app/Services/Ozon/Methods"
	"main/internal/app/Services/Ozon/Structures"
	"main/internal/app/Services/Ozon/Structures/Categories"
)

//CategoryAttribute Список характеристик категории/**
func CategoryAttribute(body Categories.AttrBody) (*Categories.AttrsResult, error) {
	resp, err := Methods.OzonResponse(body, Structures.OzonMethod{Method: "POST", Url: "/v3/category/attribute"})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result := Categories.AttrsResult{}
	err = json.NewDecoder(io.MultiReader(resp.Body)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}
