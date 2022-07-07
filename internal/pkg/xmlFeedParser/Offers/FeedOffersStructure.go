package Offers

import "encoding/xml"

type OffersDef struct {
	Offers xml.Name   `xml:"offers"`
	Offer  []OfferDef `xml:"offer"`
}

type OfferDef struct {
	Offer           []OfferDef   `xml:"offer"`
	Id              string       `xml:"id,attr"`
	OfferAttributes OfferAttrDef `xml:"attrs"`
	OfferCard       OfferCardDef `xml:"card"`
}

type OfferAttrDef struct {
	Attr []xml.Name `xml:",any"`
}

type OfferCardDef struct {
	Id string `xml:"id,attr"`
}
