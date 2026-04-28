package structures

import (
	"encoding/xml"
	"pxm/modules/filters"
	"pxm/modules/globals"
	"pxm/modules/regions"
)

type Structures struct {
	XMLName    xml.Name    `xml:"structures"`
	Structures []Structure `xml:"structure"`
	Dynamics   []Dynamic   `xml:"dynamic"`
	globals.Globals
}

type Structure struct {
	XMLName      xml.Name                 `xml:"structure"`
	ID           *string                  `xml:"id,attr,omitempty"`
	InlineRegion *string                  `xml:"region,attr,omitempty"`
	Origin       *string                  `xml:"origin,attr,omitempty"`
	Air          *string                  `xml:"air,attr,omitempty"`
	Clear        *string                  `xml:"clear,attr,omitempty"`
	Region       *regions.RegionContainer `xml:"region,omitempty"`
	globals.Globals
}

type Dynamic struct {
	XMLName       xml.Name                 `xml:"dynamic"`
	ID            *string                  `xml:"id,attr,omitempty"`
	Structure     string                   `xml:"structure,attr"`
	InlineFilter  *string                  `xml:"filter,attr,omitempty"`
	InlineTrigger *string                  `xml:"trigger,attr,omitempty"`
	Location      *string                  `xml:"location,attr,omitempty"`
	Offset        *string                  `xml:"offset,attr,omitempty"`
	Update        *string                  `xml:"update,attr,omitempty"`
	Filter        *filters.FilterContainer `xml:"filter,omitempty"`
	Trigger       *filters.FilterContainer `xml:"trigger,omitempty"`
	globals.Globals
}
