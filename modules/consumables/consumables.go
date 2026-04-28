package consumables

import (
	"encoding/xml"
	"pxm/modules/globals"
)

type Consumables struct {
	XMLName     xml.Name     `xml:"consumables"`
	Consumables []Consumable `xml:"consumable"`
	globals.Globals
}

type Consumable struct {
	XMLName  xml.Name `xml:"consumable"`
	ID       string   `xml:"id,attr"`
	Action   string   `xml:"action,attr"`
	On       string   `xml:"on,attr"`
	Override *string  `xml:"override,attr,omitempty"`
	Consume  *string  `xml:"consume,attr,omitempty"`
}
