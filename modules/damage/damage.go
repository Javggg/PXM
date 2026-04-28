package damage

import (
	"encoding/xml"
	"pxm/modules/filters"
	"pxm/modules/globals"
)

type Damage struct {
	XMLName        xml.Name `xml:"damage"`
	AttackerAction *string  `xml:"attacker-action,attr,omitempty"`
	VictimAction   *string  `xml:"victim-action,attr,omitempty"`
	*filters.FilterContainer
}

type FriendlyFire struct {
	XMLName xml.Name `xml:"friendlyfire"`
	Value   []byte   `xml:",innerxml"`
}

type FriendlyFireRefund struct {
	XMLName xml.Name `xml:"friendlyfirerefund"`
	Value   []byte   `xml:",innerxml"`
}

type Difficulty struct {
	XMLName xml.Name `xml:"difficulty"`
	Value   []byte   `xml:",innerxml"`
}

type Hunger struct {
	XMLName   xml.Name `xml:"hunger"`
	Depletion *string  `xml:"depletion"`
}

type DisableDamage struct {
	XMLName xml.Name   `xml:"disabledamage"`
	Damage  []struct { // TODO: only cause explosion
		Ally  *string `xml:"ally,attr,omitempty"`
		Self  *string `xml:"self,attr,omitempty"`
		Enemy *string `xml:"enemy,attr,omitempty"`
		Other *string `xml:"other,attr,omitempty"`
		Value []byte  `xml:",innerxml"`
	} `xml:"damage"`
	globals.Globals
}
