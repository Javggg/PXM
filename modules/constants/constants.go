package constants

import (
	"encoding/xml"
	"pxm/modules/globals"
)

type Constants struct {
	XMLName   xml.Name   `xml:"constants"`
	Constants []Constant `xml:"constant"`
	globals.Globals
}

type Constant struct {
	XMLName  xml.Name `xml:"constant"`
	ID       string   `xml:"id,attr"`
	Delete   *string  `xml:"delete,attr,omitempty"`
	Fallback *string  `xml:"fallback,attr,omitempty"`
	Value    []byte   `xml:",innerxml"`
}
