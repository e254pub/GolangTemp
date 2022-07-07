package Cards

import "encoding/xml"

type CardsDef struct {
	Cards xml.Name  `xml:"cards"`
	Card  []CardDef `xml:"card"`
}

type CardDef struct {
	Card           []CardDef   `xml:"card"`
	Id             string      `xml:"id,attr"`
	AttributesCard AttrCardDef `xml:"attrs"`
}

type AttrCardDef struct {
	Attr []xml.Name `xml:",any"`
}
