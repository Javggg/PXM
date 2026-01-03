package renewables

import (
	"encoding/xml"
	"pxm/modules/filters"
	"pxm/modules/regions"
)

type Renewables struct {
	XMLName    xml.Name    `xml:"renewables"`
	Renewables []Renewable `xml:"renewable"`
}

type Renewable struct {
	XMLName             xml.Name                 `xml:"renewable"`
	InlineRegion        *string                  `xml:"region,attr,omitempty"`
	InlineRenewFilter   *string                  `xml:"renew-filter,attr,omitempty"`
	InlineReplaceFilter *string                  `xml:"replace-filter,attr,omitempty"`
	InlineShuffleFilter *string                  `xml:"shuffle-filter,attr,omitempty"`
	Rate                *string                  `xml:"rate,attr,omitempty"`
	Interval            *string                  `xml:"interval,attr,omitempty"`
	Grow                *string                  `xml:"grow,attr,omitempty"`
	Particles           *string                  `xml:"particles,attr,omitempty"`
	Sound               *string                  `xml:"sound,attr,omitempty"`
	AvoidPlayers        *string                  `xml:"avoid-players,attr,omitempty"`
	Region              *regions.RegionContainer `xml:"region,omitempty"`
	RenewFilter         *filters.FilterContainer `xml:"renew-filter,omitempty"`
	ReplaceFilter       *filters.FilterContainer `xml:"replace-filter,omitempty"`
	ShuffleFilter       *filters.FilterContainer `xml:"shuffle-filter,omitempty"`
}
