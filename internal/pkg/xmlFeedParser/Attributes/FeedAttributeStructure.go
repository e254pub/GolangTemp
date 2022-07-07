package Attributes

import "encoding/xml"

type AttrDef struct {
	Attr       xml.Name        `xml:"attrs"`
	Attributes []AttributesDef `xml:"a"`
}

type AttributesDef struct {
	Id         string          `xml:"id,attr"`
	Type       string          `xml:"type,attr"`
	V          ValuesDef       `xml:"v"`
	Attributes []AttributesDef `xml:"a"`
}

type ValuesDef struct {
	Units string `xml:"units,attr"`
	V     string `xml:",innerxml"`
}
