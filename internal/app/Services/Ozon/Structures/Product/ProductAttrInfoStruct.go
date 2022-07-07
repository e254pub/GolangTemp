package Product

import (
	"main/internal/app/Services/Ozon"
	"main/internal/app/Services/Ozon/Structures"
	"reflect"
	"strconv"
)

type ProductInfoAttributesBody struct {
	Filter  Structures.Filter `json:"filter"`
	LastId  string            `json:"last_id"`
	Limit   int64             `json:"limit"`
	SortBy  string            `json:"sort_by"`
	SortDir string            `json:"sort_dir"`
}

type ProductsInfoAttrResult struct {
	Result []ProductInfoAttrResult `json:"result"`
	Total  int                     `json:"total"`
	LastID string                  `json:"last_id"`
}

type ProductInfoAttrResult struct {
	ID         int    `json:"id"`
	CategoryID int    `json:"category_id"`
	OfferID    string `json:"offer_id"`
	BasicPropertyProduct
	Images []struct {
		FileName string `json:"file_name"`
		Default  bool   `json:"default"`
		Index    int    `json:"index"`
	} `json:"images"`
	ImageGroupID string        `json:"image_group_id"`
	Images360    []interface{} `json:"images360"`
	PdfList      []interface{} `json:"pdf_list"`
	Attributes   []struct {
		AttributeID int `json:"attribute_id"`
		ComplexID   int `json:"complex_id"`
		Values      []struct {
			DictionaryValueID int    `json:"dictionary_value_id"`
			Value             string `json:"value"`
		} `json:"values"`
	} `json:"attributes"`
	ComplexAttributes []interface{} `json:"complex_attributes"`
	ColorImage        string        `json:"color_image"`
	LastID            string        `json:"last_id"`
}

// BasicPropertyProduct структура для ключей базовых атрибутов
type BasicPropertyProduct struct {
	Barcode       string `json:"barcode"`
	Name          string `json:"name"`
	Height        int    `json:"height"`
	Depth         int    `json:"depth"`
	Width         int    `json:"width"`
	DimensionUnit string `json:"dimension_unit"`
	Weight        int    `json:"weight"`
	WeightUnit    string `json:"weight_unit"`
}

// GetRowsProductInfo извлекает информацию
func (a *ProductInfoAttrResult) GetRowsProductInfo(r []Structures.ProductStruct, attributes []Structures.AttributeStruct) ([]Structures.ProductStruct, []Structures.AttributeStruct) {
	parentId := a.OfferID
	externalId := strconv.Itoa(a.ID)
	var dictionaryAttributes []Structures.AttributeStruct
	var productAttributes []string
	dictionaryAttributes = getDictionaryAttributes(a, dictionaryAttributes)
	dictionaryAttributes = getBasicAttributes(a, dictionaryAttributes)
	for i := 0; i < len(dictionaryAttributes); i++ {
		productAttributes = append(productAttributes, dictionaryAttributes[i].ExternalId)
	}
	row := Structures.ProductStruct{"набор полей"}
	r = append(r, row)
	attributes = append(attributes, dictionaryAttributes...)
	return r, attributes
}

// getDictionaryAttributes Пополняет массив словарными атрибутами из товара
func getDictionaryAttributes(a *ProductInfoAttrResult, dictionaryAttributes []Structures.AttributeStruct) []Structures.AttributeStruct {
	categoryId := Ozon.CATEGORY + "_" + strconv.Itoa(a.CategoryID)
	productId := strconv.Itoa(a.ID)
	for i := 0; i < len(a.Attributes); i++ {
		attributeID := strconv.Itoa(a.Attributes[i].AttributeID)
		for _, value := range a.Attributes[i].Values {
			if value.Value != "" {
				dictionaryId := Ozon.DICTIONARY_ATTRIBUTE + "_" + strconv.Itoa(value.DictionaryValueID)
				dictionaryValue := value.Value
				if value.DictionaryValueID == 0 {
					dictionaryId = Ozon.ATTRIBUTE + "_" + attributeID + "_" + productId
					attributeID = categoryId
				}
				rowAttribute := Structures.AttributeStruct{
					"набор полей"}
				dictionaryAttributes = append(dictionaryAttributes, rowAttribute)
			}
		}
	}
	return dictionaryAttributes
}

// getBasicAttributes пополняет базовыми атрибутами из товара
func getBasicAttributes(a *ProductInfoAttrResult, dictionaryAttributes []Structures.AttributeStruct) []Structures.AttributeStruct {
	var b BasicPropertyProduct
	val := reflect.ValueOf(b)
	t := val.Type()
	var valueProperties []string
	valueProperties = append(valueProperties, a.Barcode, a.Name, strconv.Itoa(a.Height), strconv.Itoa(a.Depth),
		strconv.Itoa(a.Width), a.DimensionUnit, strconv.Itoa(a.Weight), a.WeightUnit)
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Tag.Get("json")
		rowAttribute := Structures.AttributeStruct{
			"набор полей"}
		dictionaryAttributes = append(dictionaryAttributes, rowAttribute)
	}
	return dictionaryAttributes
}
