package teams

import "encoding/xml"

type Teams struct {
	XMLName xml.Name `xml:"teams"`
	Teams   []Team   `xml:"team"`
}

type Team struct {
	XMLName      xml.Name `xml:"team"`
	ID           string   `xml:"id,attr"`
	Color        *string  `xml:"color,attr,omitempty"`
	DyeColor     *string  `xml:"dye-color,attr,omitempty"`
	Plural       *string  `xml:"plural,attr,omitempty"`
	ShowNameTags *string  `xml:"show-name-tags,attr,omitempty"`
	Min          *string  `xml:"min,attr,omitempty"`
	Max          string   `xml:"max,attr,omitempty"`
	MaxOverfill  *string  `xml:"max-overfill,attr,omitempty"`
	Name         string   `xml:",innerxml"`
}

type Players struct {
	Min          *string `xml:"min,attr,omitempty"`
	Max          string  `xml:"max,attr,omitempty"`
	MaxOverfill  *string `xml:"max-overfill,attr,omitempty"`
	ShowNameTags *string `xml:"show-name-tags,attr,omitempty"`
	Colors       *string `xml:"colors,attr,omitempty"`
}
