package score

import (
	"encoding/xml"
	"pxm/modules/filters"
	"pxm/modules/regions"
)

type Score struct {
	XMLName          xml.Name `xml:"score"`
	Display          *string  `xml:"display,attr,omitempty"`
	ScoreboardFilter *string  `xml:"scoreboard-filter,attr,omitempty"`
	Limit            *string  `xml:"limit,omitempty"`
	Mercy            *string  `xml:"mercy,omitempty"`
	Kills            *string  `xml:"kills,omitempty"`
	Deaths           *string  `xml:"deaths,omitempty"`
	Boxes            []Box    `xml:"box"`
}

type Box struct {
	XMLName       xml.Name                 `xml:"box"`
	Points        *string                  `xml:"points,attr,omitempty"`
	Silent        *string                  `xml:"silent,attr,omitempty"`
	InlineRegion  *string                  `xml:"region,attr,omitempty"`
	InlineFilter  *string                  `xml:"filter,attr,omitempty"`
	InlineTrigger *string                  `xml:"trigger,attr,omitempty"`
	Region        *regions.RegionContainer `xml:"region,omitempty"`
	Filter        *filters.FilterContainer `xml:"filter,omitempty"`
	Trigger       *filters.FilterContainer `xml:"trigger,omitempty"`
	Redeemables   []struct {
		Items []struct {
			Points *string `xml:"points,attr,omitempty"`
			Value  string  `xml:",chardata"`
		} `xml:"item"`
	} `xml:"redeemables"`
}

type Time struct {
	XMLName     xml.Name `xml:"time"`
	Result      *string  `xml:"result,attr,omitempty"`
	Show        *string  `xml:"show,attr,omitempty"`
	Overtime    *string  `xml:"overtime,attr,omitempty"`
	MaxOvertime *string  `xml:"max-overtime,attr,omitempty"`
	Value       *string  `xml:",chardata"`
}
