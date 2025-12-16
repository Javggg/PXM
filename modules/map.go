package modules

import "encoding/xml"

type Map struct {
	XMLName          xml.Name `xml:"map"`
	Proto            string   `xml:"proto,attr"`
	MinServerVersion *string  `xml:"min-server-version,attr"`
	MaxServerVersion *string  `xml:"max-server-version,attr"`
	Internal         *bool    `xml:"internal,attr"`
	Authors          *Authors `xml:"authors,omitempty"`
	Regions          *Regions `xml:"regions,omitempty"`
}
