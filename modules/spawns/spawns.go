package spawns

import (
	"encoding/xml"
	"pxm/modules/kits"
	"pxm/modules/regions"
)

type Spawns struct {
	XMLName   xml.Name  `xml:"spawns"`
	Default   Spawn     `xml:"default"`
	Spawns    []Spawn   `xml:"spawn"`
	PlayerKit *kits.Kit `xml:"player-kit"`
}

type Spawn struct {
	Team            *string   `xml:"team,attr"`
	InlineRegion    *string   `xml:"region,attr,omitempty"`
	Regions         []Regions `xml:"regions"`
	Safe            *string   `xml:"safe,attr,omitempty"`
	Sequential      *string   `xml:"sequential,attr,omitempty"`
	Spread          *string   `xml:"spread,attr,omitempty"`
	SpreadTeammates *string   `xml:"spread-teammates,attr,omitempty"`
	Exclusive       *string   `xml:"exclusivem,attr,omitempty"`
	Outdoors        *string   `xml:"outdoors,attr,omitempty"`
	Persistent      *string   `xml:"persistent,attr,omitempty"`
	Kit             *string   `xml:"kit,attr,omitempty"`
	Filter          *string   `xml:"filter,attr,omitempty"`
}

type Regions struct {
	XMLName xml.Name `xml:"regions"`
	Angle   *string  `xml:"angle,attr,omitempty"`
	Yaw     *string  `xml:"yaw,attr,omitempty"`
	Pitch   *string  `xml:"pitch,attr,omitempty"`
	regions.RegionContainer
}
