package Database

import (
	"fmt"
	Categories2 "main/internal/app/Services/Ozon/Methods/Categories"
	"main/internal/app/Services/Ozon/Structures"
	"main/internal/app/Services/Ozon/Structures/Categories"
	"time"
)

// GetAttributesSlice возвращает записи на базу с уникальными аттрибутами
func GetAttributesSlice(CategoryIds []int) (int64, error) {
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})
	var attrIds []Categories.AttrIds
	var records []Structures.AttributeStruct
	for i := 0; i < len(CategoryIds); i += 20 {
		select {
		case <-ticker.C:
			CategoryIdsSlice := CategoryIds[i : i+20]
			body := Categories.AttrBody{
				CategoryID:    CategoryIdsSlice,
				Language:      "DEFAULT",
				AttributeType: "ALL",
			}
			res, _ := Categories2.CategoryAttribute(body)
			attrUnsort := attributesSlice(res)
			fmt.Println(i)
			attrIds = append(attrIds, attrUnsort...)
		case <-quit:
			ticker.Stop()
		}
	}
	sortAttr := ozonUniqueAttrs(attrIds)

	for _, value := range sortAttr {
		records = value.GetRowsAttributes(records)
	}
	return AttributeRecords(records)
}

func GetAttributeIds(ShopId int, Limit interface{}) []int {
	Query := ""
	rowSlice := GetDbAttrs(Query, ShopId, Limit)
	return rowSlice
}

// attributesSlice возвращает список неотсортированных id аттрибутов
func attributesSlice(result *Categories.AttrsResult) []Categories.AttrIds {
	catalogRes := result.Result
	var attrIds []Categories.AttrIds
	for i := 0; i < len(catalogRes); i++ {
		attrRes := catalogRes[i].Attributes
		for _, value := range attrRes {
			attrIds = value.GetAttributesIds(attrIds)
		}
	}
	return attrIds
}

// ozonUniqueAttrs возвращает уникальные значения
func ozonUniqueAttrs(sample []Categories.AttrIds) []Categories.AttrIds {
	var unique []Categories.AttrIds
sampleLoop:
	for _, v := range sample {
		for i, u := range unique {
			if "набор полей" {
				unique[i] = v
				continue sampleLoop
			}
		}
		unique = append(unique, v)
	}
	return unique
}
