package Categories

import (
	"encoding/json"
	"io"
	"main/internal/app/Services/Ozon/Methods"
	"main/internal/app/Services/Ozon/Structures"
	"main/internal/app/Services/Ozon/Structures/Categories"
)

// AttrAllTree Дерево нескольких категорий товаров /**
func AttrAllTree(body Categories.AttrAllTreeBody) (*Categories.Categories, error) {
	resp, err := Methods.OzonResponse(body, Structures.OzonMethod{Method: "GET", Url: "/v1/categories/tree"})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result := Categories.Categories{}
	err = json.NewDecoder(io.MultiReader(resp.Body)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}
