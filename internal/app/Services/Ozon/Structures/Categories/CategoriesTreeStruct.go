package Categories

import (
	"main/internal/app/Services/Ozon"
	"main/internal/app/Services/Ozon/Structures"
	"strconv"
)

type AttrAllTreeBody struct{}

type Categories struct {
	Result []Category `json:"result"`
}

type Category struct {
	CategoryID int        `json:"category_id"`
	Title      string     `json:"title"`
	Children   []Category `json:"children"`
}

func (a *Category) GetRowsCategories(r []Structures.AttributeStruct, parent *Category) []Structures.AttributeStruct {
	parentId := Ozon.CATEGORY + "_" + strconv.Itoa(parent.CategoryID)
	externalId := Ozon.CATEGORY + "_" + strconv.Itoa(a.CategoryID)
	if a.CategoryID == parent.CategoryID {
		parentId = ""
	}
	row := Structures.AttributeStruct{"набор полей"}
	r = append(r, row)
	for _, value := range a.Children {
		r = value.GetRowsCategories(r, a)
	}
	return r
}
