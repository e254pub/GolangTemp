package Categories

import (
	"encoding/json"
	"io"
	"main/internal/app/Services/Ozon/Methods"
	"main/internal/app/Services/Ozon/Structures"
	"main/internal/app/Services/Ozon/Structures/Categories"
)

// AttrOneTree Дерево категории товаров/**
func AttrOneTree(body Categories.AttrOneTreeBody) (*Categories.AttrOneTreeResult, error) {
	resp, err := Methods.OzonResponse(body, Structures.OzonMethod{Method: "POST", Url: "/v2/category/tree"})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result := Categories.AttrOneTreeResult{}
	err = json.NewDecoder(io.MultiReader(resp.Body)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}
