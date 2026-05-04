package compass

import (
	"encoding/xml"
	"pxm/modules/globals"
)

type Compass struct {
	XMLName      xml.Name `xml:"compass"`
	Order        *string  `xml:"order,attr,omitempty"`
	ShowDistance *string  `xml:"show-distance,attr,omitempty"`
	Players      []Player `xml:"player"`
	Flags        []Flag   `xml:"flag"`
	globals.Globals
}

type Player struct {
	XMLName      xml.Name `xml:"player"`
	Filter       *string  `xml:"filter,attr,omitempty"`
	HolderFilter *string  `xml:"holder-filter,attr,omitempty"`
	Name         *string  `xml:"name,attr,omitempty"`
	ShowPlayer   *string  `xml:"show-player,attr,omitempty"`
}

type Flag struct {
	XMLName      xml.Name `xml:"flag"`
	Filter       *string  `xml:"filter,attr,omitempty"`
	HolderFilter *string  `xml:"holder-filter,attr,omitempty"`
	Name         *string  `xml:"name,attr,omitempty"`
	Value        []byte   `xml:",innerxml"`
}
