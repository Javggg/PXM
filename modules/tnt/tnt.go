package tnt

import (
	"encoding/xml"
	"pxm/modules/globals"
)

type TNT struct {
	XMLName                xml.Name `xml:"tnt"`
	InstantIgnite          *string  `xml:"instantignite,omitempty"`
	BlockDamage            *string  `xml:"blockdamage,omitempty"`
	Yield                  *string  `xml:"yield,omitempty"`
	Power                  *string  `xml:"power,omitempty"`
	Fuse                   *string  `xml:"fuse,omitempty"`
	DispenserTntLimit      *string  `xml:"dispenser-tnt-limit,omitempty"`
	DispenserTntMultiplier *string  `xml:"dispenser-tnt-multiplier,omitempty"`
	Licensing              *string  `xml:"licensing,omitempty"`
	FriendlyDefuse         *string  `xml:"friendly-defuse,omitempty"`
	globals.Globals
}
