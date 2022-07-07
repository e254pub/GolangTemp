package xmlFeedParser

import (
	"encoding/xml"
	"main/internal/pkg/xmlFeedParser/Attributes"
	Cards2 "main/internal/pkg/xmlFeedParser/Cards"
	Offers2 "main/internal/pkg/xmlFeedParser/Offers"
)

type CatalogDef struct {
	Catalog xml.Name `xml:"yml_catalog"`
	Shop    ShopDef  `xml:"shop"`
}

type ShopDef struct {
	Shop   xml.Name           `xml:"shop"`
	Attrs  Attributes.AttrDef `xml:"attrs"`
	Cards  Cards2.CardsDef    `xml:"cards"`
	Offers Offers2.OffersDef  `xml:"offers"`
}
