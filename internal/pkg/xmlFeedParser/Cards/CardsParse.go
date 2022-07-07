package Cards

import (
	"context"
	"main/internal/pkg/xmlFeedParser"
)

func CardsParse(c xmlFeedParser.CatalogDef, err error, conn *pgx.Conn) error {
	xmlResult := c.Shop.Cards.Card
	var records []xmlFeedParser.CardsRow
	for _, value := range xmlResult {
		records = value.getRowsCard(records, &value)
	}

	_, err = conn.CopyFrom(context.Background(), pgx.Identifier{"схема", "таблица"},
		[]string{"набор полей"},
		pgx.CopyFromSlice(len(records), func(i int) ([]interface{}, error) {
			return []interface{}{"набор полей"}, nil
		}),
	)

	return err
}

func (a *CardDef) getRowsCard(r []xmlFeedParser.CardsRow, parent *CardDef) []xmlFeedParser.CardsRow {
	var idParent string
	if len(parent.Card) > 0 {
		idParent = parent.Card[0].Id
	}
	var attr []string
	for i := 0; i < len(parent.AttributesCard.Attr); i++ {
		attr = append(attr, parent.AttributesCard.Attr[i].Local)
	}
	row := xmlFeedParser.CardsRow{"набор полей"}
	r = append(r, row)
	return r
}
