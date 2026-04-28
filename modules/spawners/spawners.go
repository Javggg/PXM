package spawners

import (
	"encoding/xml"
	"pxm/modules/globals"
	"pxm/modules/kits"
)

type Spawners struct {
	XMLName      xml.Name  `xml:"spawners"`
	SpawnRegion  *string   `xml:"spawn-region,attr"`
	PlayerRegion *string   `xml:"player-region,attr"`
	MaxEntities  *string   `xml:"max-entities,attr,omitempty"`
	Delay        *string   `xml:"delay,attr,omitempty"`
	MinDelay     *string   `xml:"min-delay,attr,omitempty"`
	MaxDelay     *string   `xml:"max-delay,attr,omitempty"`
	Filter       *string   `xml:"filter,attr,omitempty"`
	Spawners     []Spawner `xml:"spawner"`
	globals.Globals
}

type Spawner struct {
	XMLName      xml.Name   `xml:"spawner"`
	SpawnRegion  *string    `xml:"spawn-region,attr"`
	PlayerRegion *string    `xml:"player-region,attr"`
	MaxEntities  *string    `xml:"max-entities,attr,omitempty"`
	Delay        *string    `xml:"delay,attr,omitempty"`
	MinDelay     *string    `xml:"min-delay,attr,omitempty"`
	MaxDelay     *string    `xml:"max-delay,attr,omitempty"`
	Filter       *string    `xml:"filter,attr,omitempty"`
	Item         *kits.Item `xml:"item,omitempty"`
	Potion       *struct {
		Effects []kits.Effect `xml:"effect"`
		globals.Globals
	} `xml:"potion,omitempty"`
	globals.Globals
}
