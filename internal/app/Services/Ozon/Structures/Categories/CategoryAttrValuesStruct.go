package Categories

import (
	"main/internal/app/Services/Ozon"
	"main/internal/app/Services/Ozon/Structures"
	"strconv"
)

type AttrChildTreeBody struct {
	AttributeID int    `json:"attribute_id"`
	CategoryID  int    `json:"category_id"`
	Language    string `json:"language"`
	LastValueID int    `json:"last_value_id"`
	Limit       int    `json:"limit"`
}

type AttrDictionaryResult struct {
	Result  []AttrDictionary `json:"result"`
	HasNext bool             `json:"has_next"`
}

type AttrDictionary struct {
	Id      int    `json:"id"`
	Info    string `json:"info"`
	Picture string `json:"picture"`
	Value   string `json:"value"`
}

func (a *AttrDictionary) GetRowsDictionary(r []Structures.AttributeStruct, attrId int) []Structures.AttributeStruct {
	parentId := Ozon.ATTRIBUTE + "_" + strconv.Itoa(attrId)
	externalId := Ozon.DICTIONARY_ATTRIBUTE + "_" + strconv.Itoa(a.Id)
	row := Structures.AttributeStruct{"набор полей"}
	r = append(r, row)
	return r
}
