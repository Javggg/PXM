package replacements

import (
	"encoding/xml"
	"pxm/modules/filters"
)

type Replacements struct {
	XMLName xml.Name `xml:"replacements"`
	ReplacementsContainer
}

type ReplacementsContainer struct {
	Decimal     []Decimal     `xml:"decimal"`
	Player      []Player      `xml:"player"`
	Switch      []Switch      `xml:"switch"`
	Replacement []Replacement `xml:"replacement"`
}

type BaseReplacement struct {
	ID string `xml:"id,attr"`
}

type Replacement struct {
	ID       string `xml:"id,attr"`
	GlobalID string `xml:",chardata"`
}

type Decimal struct {
	XMLName xml.Name `xml:"decimal"`
	BaseReplacement
	Value  string  `xml:"value,attr"`
	Format *string `xml:"format,attr,omitempty"`
	Scope  *string `xml:"scope,attr,omitempty"`
}

type Player struct {
	XMLName xml.Name `xml:"player"`
	BaseReplacement
	Var      string  `xml:"var,attr"`
	Style    *string `xml:"style,attr,omitempty"`
	Fallback *string `xml:"fallback,omitempty"`
}

type Switch struct {
	XMLName xml.Name `xml:"switch"`
	BaseReplacement
	Value    string  `xml:"value,attr"`
	Scope    *string `xml:"scope,attr,omitempty"`
	Cases    []Case  `xml:"case"`
	Fallback *string `xml:"fallback,omitempty"`
}

type Case struct {
	XMLName      xml.Name                 `xml:"case"`
	InlineMatch  *string                  `xml:"match,attr,omitempty"`
	InlineFilter *string                  `xml:"filter,attr,omitempty"`
	InlineResult *string                  `xml:"result,attr,omitempty"`
	Match        *string                  `xml:"match,omitempty"`
	Filter       *filters.FilterContainer `xml:"filter,omitempty"`
	Result       *string                  `xml:"result,omitempty"`
}
