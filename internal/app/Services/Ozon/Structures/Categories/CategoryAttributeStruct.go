package Categories

import (
	"main/internal/app/Services/Ozon"
	"main/internal/app/Services/Ozon/Structures"
	"strconv"
)

type AttrBody struct {
	AttributeType string `json:"attribute_type"`
	CategoryID    []int  `json:"category_id"`
	Language      string `json:"language"`
}

type AttrsResult struct {
	Result []Attr `json:"result"`
}

type Attr struct {
	Attributes []Attributes `json:"attributes"`
	CategoryID int          `json:"category_id"`
}

type Attributes struct {
	Description  string `json:"description"`
	DictionaryId int    `json:"dictionary_id"`
	GroupId      int    `json:"group_id"`
	GroupName    string `json:"group_name"`
	Id           int    `json:"id"`
	IsCollection bool   `json:"is_collection"`
	IsRequired   bool   `json:"is_required"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

type AttrIds struct {
	ExternalId int
	Value      string
}

func (a *Attributes) GetAttributesIds(r []AttrIds) []AttrIds {
	externalId := a.Id
	row := AttrIds{externalId, a.Name}
	r = append(r, row)
	return r
}

func (a *AttrIds) GetRowsAttributes(r []Structures.AttributeStruct) []Structures.AttributeStruct {
	externalId := Ozon.ATTRIBUTE + "_" + strconv.Itoa(a.ExternalId)
	row := Structures.AttributeStruct{"набор полей"}
	r = append(r, row)
	return r
}
