package modes

import (
	"encoding/xml"
	"pxm/modules/globals"
)

type Modes struct {
	XMLName xml.Name `xml:"modes"`
	Modes   []Mode   `xml:"mode"`
	globals.Globals
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
