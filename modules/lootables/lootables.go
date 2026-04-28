package lootables

import (
	"pxm/modules/globals"
	"pxm/modules/kits"
)

type Lootables struct {
	Loots []Loot `xml:"loot"`
	Fills []Fill `xml:"fill"`
	globals.Globals
}

type LootContainer struct {
	Anys    []Any       `xml:"any"`
	Maybes  []Maybe     `xml:"maybe"`
	Alls    []All       `xml:"all"`
	Items   []kits.Item `xml:"item"`
	Options []Option    `xml:"option"`
	globals.Globals
}

type Loot struct {
	ID *string `xml:"id,attr,omitempty"`
	LootContainer
}

type Any struct {
	Count  *string `xml:"count,attr,omitempty"`
	Unique *string `xml:"unique,attr,omitempty"`
	LootContainer
}

type Maybe struct {
	Filter *string     `xml:"filter,attr,omitempty"`
	Items  []kits.Item `xml:"item"`
	LootContainer
}

type All struct {
	LootContainer
}

type Option struct {
	Weight *string `xml:"weight,attr,omitempty"`
	Filter *string `xml:"filter,attr,omitempty"`
	LootContainer
}

type Fill struct {
	Loot           string  `xml:"loot,attr"`
	Filter         *string `xml:"filter,attr,omitempty"`
	RefillTrigger  *string `xml:"refill-trigger,attr,omitempty"`
	RefillInterval *string `xml:"refill-interval,attr,omitempty"`
	RefillClear    *string `xml:"refill-clear,attr,omitempty"`
}
