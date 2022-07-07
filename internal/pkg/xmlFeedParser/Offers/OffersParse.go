package Offers

import (
	"context"
	"main/internal/pkg/xmlFeedParser"
)

func OffersParse(c xmlFeedParser.CatalogDef, err error, conn *pgx.Conn) error {
	xmlResult := c.Shop.Offers.Offer
	var records []xmlFeedParser.OffersRow
	for _, value := range xmlResult {
		records = value.getRowsOffers(records, &value)
	}

	_, err = conn.CopyFrom(context.Background(), pgx.Identifier{"схема", "таблица"},
		[]string{"набор полей"},
		pgx.CopyFromSlice(len(records), func(i int) ([]interface{}, error) {
			return []interface{}{"набор полей"}, nil
		}),
	)

	return err
}

func (a *OfferDef) getRowsOffers(r []xmlFeedParser.OffersRow, parent *OfferDef) []xmlFeedParser.OffersRow {
	var attr []string
	for i := 0; i < len(parent.OfferAttributes.Attr); i++ {
		attr = append(attr, parent.OfferAttributes.Attr[i].Local)
	}
	row := xmlFeedParser.OffersRow{"набор полей"}
	r = append(r, row)
	return r
}
