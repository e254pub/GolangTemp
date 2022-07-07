package Categories

import (
	"encoding/json"
	"io"
	"main/internal/app/Services/Ozon/Methods"
	"main/internal/app/Services/Ozon/Structures"
	"main/internal/app/Services/Ozon/Structures/Categories"
)

// DictionaryAttr Справочник значений характеристики/**
func DictionaryAttr(body Categories.AttrChildTreeBody) (*Categories.AttrDictionaryResult, error) {
	resp, err := Methods.OzonResponse(body, Structures.OzonMethod{Method: "POST", Url: "/v2/category/attribute/values"})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	result := Categories.AttrDictionaryResult{}
	err = json.NewDecoder(io.MultiReader(resp.Body)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}
