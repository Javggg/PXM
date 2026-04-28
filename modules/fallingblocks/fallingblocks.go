package fallingblocks

import (
	"encoding/xml"
	"pxm/modules/filters"
	"pxm/modules/globals"
)

type FallingBlocks struct {
	XMLName xml.Name `xml:"falling-blocks"`
	Rules   []Rule   `xml:"rule"`
	globals.Globals
}

type Rule struct {
	InlineFilter *string                  `xml:"filter,attr,omitempty"`
	InlineSticky *string                  `xml:"sticky,attr,omitempty"`
	InlineDelay  *string                  `xml:"delay,attr,omitempty"`
	Filter       *filters.FilterContainer `xml:"filter,omitempty"`
	Sticky       *filters.FilterContainer `xml:"sticky,omitempty"`
	Delay        *string                  `xml:"delay,omitempty"`
	globals.Globals
}
