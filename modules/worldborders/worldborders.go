package worldborders

import (
	"encoding/xml"
	"pxm/modules/filters"
)

type WorldBorders struct {
	XMLName      xml.Name      `xml:"world-borders"`
	WorldBorders []WorldBorder `xml:"world-border"`
}

type WorldBorder struct {
	XMLName         xml.Name                 `xml:"world-border"`
	Center          string                   `xml:"center,attr"`
	Size            string                   `xml:"size,attr"`
	InlineWhen      *string                  `xml:"when,attr,omitempty"`
	After           *string                  `xml:"after,attr,omitempty"`
	Duration        *string                  `xml:"duration,attr,omitempty"`
	Damage          *string                  `xml:"damage,attr,omitempty"`
	Buffer          *string                  `xml:"buffer,attr,omitempty"`
	WarningDistance *string                  `xml:"warning-distance,attr,omitempty"`
	WarningTime     *string                  `xml:"warning-time,attr,omitempty"`
	When            *filters.FilterContainer `xml:"when,omitempty"`
}
