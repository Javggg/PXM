package portals

import (
	"encoding/xml"
	"pxm/modules/filters"
	"pxm/modules/globals"
	"pxm/modules/regions"
)

type Portals struct {
	XMLName       xml.Name `xml:"portals"`
	Filter        *string  `xml:"filter,attr,omitempty"`
	Sound         *string  `xml:"sound,attr,omitempty"`
	Protect       *string  `xml:"protect,attr,omitempty"`
	Bidirectional *string  `xml:"bidirectional,attr,omitempty"`
	Smooth        *string  `xml:"smooth,attr,omitempty"`
	Yaw           *string  `xml:"yaw,attr,omitempty"`
	Pitch         *string  `xml:"pitch,attr,omitempty"`
	Observers     *string  `xml:"observers,attr,omitempty"`
	Portals       []Portal `xml:"portal"`
	globals.Globals
}

type Portal struct {
	XMLName           xml.Name                 `xml:"portal"`
	X                 *string                  `xml:"x,attr,omitempty"`
	Y                 *string                  `xml:"y,attr,omitempty"`
	Z                 *string                  `xml:"z,attr,omitempty"`
	InlineRegion      *string                  `xml:"region,attr,omitempty"`
	InlineDestination *string                  `xml:"destination,attr,omitempty"`
	InlineFilter      *string                  `xml:"filter,attr,omitempty"`
	Region            *regions.RegionContainer `xml:"region,omitempty"`
	Destination       *regions.RegionContainer `xml:"destination,omitempty"`
	Filter            *filters.FilterContainer `xml:"filter,omitempty"`
	Sound             *string                  `xml:"sound,attr,omitempty"`
	Protect           *string                  `xml:"protect,attr,omitempty"`
	Bidirectional     *string                  `xml:"bidirectional,attr,omitempty"`
	Smooth            *string                  `xml:"smooth,attr,omitempty"`
	Yaw               *string                  `xml:"yaw,attr,omitempty"`
	Pitch             *string                  `xml:"pitch,attr,omitempty"`
	InlineForward     *string                  `xml:"forward,attr,omitempty"`
	InlineReverse     *string                  `xml:"reverse,attr,omitempty"`
	InlineTransit     *string                  `xml:"transit,attr,omitempty"`
	InlineObservers   *string                  `xml:"observers,attr,omitempty"`
	Forward           *filters.FilterContainer `xml:"forward,omitempty"`
	Reverse           *filters.FilterContainer `xml:"reverse,omitempty"`
	Transit           *filters.FilterContainer `xml:"transit,omitempty"`
	Observers         *filters.FilterContainer `xml:"observers,omitempty"`
	globals.Globals
}
