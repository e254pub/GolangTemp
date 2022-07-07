package Attributes

import (
	"context"
	"main/internal/pkg/xmlFeedParser"
)

func AttributeParse(c xmlFeedParser.CatalogDef, err error, conn *pgx.Conn) error {
	xmlResult := c.Shop.Attrs.Attributes
	var records []xmlFeedParser.AttributesRow
	for _, value := range xmlResult {
		records = value.getRowsAttribute(records, &value)
	}

	_, err = conn.CopyFrom(context.Background(), pgx.Identifier{"схема", "таблица"},
		[]string{"набор полей"},
		pgx.CopyFromSlice(len(records), func(i int) ([]interface{}, error) {
			return []interface{}{"набор полей"}, nil
		}),
	)

	return err
}

func (a *AttributesDef) getRowsAttribute(r []xmlFeedParser.AttributesRow, parent *AttributesDef) []xmlFeedParser.AttributesRow {
	parentId := parent.Id
	if a.Id == parent.Id {
		parentId = ""
	}
	row := xmlFeedParser.AttributesRow{"набор полей"}
	r = append(r, row)
	for _, value := range a.Attributes {
		r = value.getRowsAttribute(r, a)
	}
	return r
}
