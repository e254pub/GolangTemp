package Database

import (
	"main/internal/app/Services/Ozon/Structures"
	"main/internal/app/Services/Ozon/Structures/Product"
)

// ProductInfoAttributesCopy расчленяет фид со списком товаров на товары и атрибуты (библиотечные и основные свойства)
func ProductInfoAttributesCopy(result *Product.ProductsInfoAttrResult) {
	productInfoRes := result.Result
	var records []Structures.ProductStruct
	var productAttributes []Structures.AttributeStruct
	for i := 0; i < len(productInfoRes); i++ {
		card := Structures.ProductStruct{
			"набор полей"
		}
		records = append(records, card)
	}
	for _, value := range productInfoRes {
		records, productAttributes = value.GetRowsProductInfo(records, productAttributes)
	}

	AttributeRecords(UniqueAttrs(productAttributes))
	ProductRecords(records)
}
