package projectiles

import (
	"encoding/xml"
	"pxm/modules/filters"
	"pxm/modules/globals"
)

type Projectiles struct {
	XMLName     xml.Name     `xml:"projectiles"`
	Projectiles []Projectile `xml:"projectile"`
	globals.Globals
}

type Projectile struct {
	XMLName             xml.Name                 `xml:"projectile"`
	ID                  string                   `xml:"id,attr"`
	Name                *string                  `xml:"name,attr,omitempty"`
	Throwable           *string                  `xml:"throwable,attr,omitempty"`
	Projectile          *string                  `xml:"projectile,attr,omitempty"`
	Power               *string                  `xml:"power,attr,omitempty"`
	Damage              *string                  `xml:"damage,attr,omitempty"`
	Velocity            *string                  `xml:"velocity,attr,omitempty"`
	Click               *string                  `xml:"click,attr,omitempty"`
	Effect              *string                  `xml:"effect,attr,omitempty"`
	InlineDestroyFilter *string                  `xml:"destroy-filter,attr,omitempty"`
	Cooldown            *string                  `xml:"cooldown,attr,omitempty"`
	Precise             *string                  `xml:"precise,attr,omitempty"`
	DestroyFilter       *filters.FilterContainer `xml:"destroy-filter,omitempty"`
	globals.Globals
}
