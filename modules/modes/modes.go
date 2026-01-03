package modes

import "encoding/xml"

type Modes struct {
	XMLName xml.Name `xml:"modes"`
	Modes   []Mode   `xml:"mode"`
}

type Mode struct {
	XMLName    xml.Name `xml:"mode"`
	ID         *string  `xml:"id,attr,omitempty"`
	Name       *string  `xml:"name,attr,omitempty"`
	Filter     *string  `xml:"filter,attr,omitempty"`
	Action     *string  `xml:"action,attr,omitempty"`
	Material   *string  `xml:"material,attr,omitempty"`
	After      string   `xml:"after,attr"`
	ShowBefore *string  `xml:"show-before,attr,omitempty"`
}
