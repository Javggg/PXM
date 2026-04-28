package wools

import (
	"encoding/xml"
	"pxm/modules/globals"
	"pxm/modules/regions"
)

type Wools struct {
	XMLName                     xml.Name `xml:"wools"`
	Required                    *string  `xml:"required,attr,omitempty"`
	Team                        *string  `xml:"team,attr,omitempty"`
	InlineMonument              *string  `xml:"monument,attr,omitempty"`
	Craftable                   *string  `xml:"craftable,attr,omitempty"`
	ShowMessages                *string  `xml:"show-messages,attr,omitempty"`
	ShowEffects                 *string  `xml:"show-effects,attr,omitempty"`
	ShowInfo                    *string  `xml:"show-info,attr,omitempty"`
	ShowSidebar                 *string  `xml:"show-sidebar,attr,omitempty"`
	Stats                       *string  `xml:"stats,attr,omitempty"`
	Show                        *string  `xml:"show,attr,omitempty"`
	ScoreboardFilter            *string  `xml:"scoreboard-filter,attr,omitempty"`
	Location                    *string  `xml:"location,attr,omitempty"`
	WoolProximityMetric         *string  `xml:"wool-proximity-metric,attr,omitempty"`
	WoolProximityHorizontal     *string  `xml:"wool-proximity-horizontal,attr,omitempty"`
	MonumentProximityMetric     *string  `xml:"monument-proximity-metric,attr,omitempty"`
	MonumentProximityHorizontal *string  `xml:"monument-proximity-horizontal,attr,omitempty"`
	Wools                       []Wool   `xml:"wool"`
	globals.Globals
}

type Wool struct {
	XMLName                     xml.Name                 `xml:"wool"`
	ID                          *string                  `xml:"id,attr,omitempty"`
	Required                    *string                  `xml:"required,attr,omitempty"`
	Team                        *string                  `xml:"team,attr,omitempty"`
	Color                       *string                  `xml:"color,attr,omitempty"`
	InlineMonument              *string                  `xml:"monument,attr,omitempty"`
	Craftable                   *string                  `xml:"craftable,attr,omitempty"`
	ShowMessages                *string                  `xml:"show-messages,attr,omitempty"`
	ShowEffects                 *string                  `xml:"show-effects,attr,omitempty"`
	ShowInfo                    *string                  `xml:"show-info,attr,omitempty"`
	ShowSidebar                 *string                  `xml:"show-sidebar,attr,omitempty"`
	Stats                       *string                  `xml:"stats,attr,omitempty"`
	Show                        *string                  `xml:"show,attr,omitempty"`
	ScoreboardFilter            *string                  `xml:"scoreboard-filter,attr,omitempty"`
	Location                    *string                  `xml:"location,attr,omitempty"`
	WoolProximityMetric         *string                  `xml:"wool-proximity-metric,attr,omitempty"`
	WoolProximityHorizontal     *string                  `xml:"wool-proximity-horizontal,attr,omitempty"`
	MonumentProximityMetric     *string                  `xml:"monument-proximity-metric,attr,omitempty"`
	MonumentProximityHorizontal *string                  `xml:"monument-proximity-horizontal,attr,omitempty"`
	Monument                    *regions.RegionContainer `xml:"monument,omitempty"`
	globals.Globals
}
