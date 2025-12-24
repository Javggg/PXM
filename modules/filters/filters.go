package filters

import (
	"encoding/xml"
	"pxm/modules/regions"
)

type FilterElement interface{}

type Filters struct {
	XMLName xml.Name        `xml:"filters"`
	Items   []FilterElement `xml:",any"`
}

type FilterRef struct {
	XMLName xml.Name `xml:"filter"`
	RefID   string   `xml:"id,attr"`
}

type BaseFilter struct {
	ID *string `xml:"id,attr,omitempty"`
}

// Generic

type Always struct {
	XMLName xml.Name `xml:"always"`
	BaseFilter
}

type Never struct {
	XMLName xml.Name `xml:"never"`
	BaseFilter
}

type Time struct {
	XMLName xml.Name `xml:"time"`
	BaseFilter
	Value string `xml:",chardata"` // check for valid duration
}

type MatchStarted struct {
	XMLName xml.Name `xml:"match-started"`
	BaseFilter
}

type MatchRunning struct {
	XMLName xml.Name `xml:"match-running"`
	BaseFilter
}

type MatchFinished struct {
	XMLName xml.Name `xml:"match-finished"`
	BaseFilter
}

type Objective struct {
	XMLName xml.Name `xml:"objective"`
	BaseFilter
	ObjectiveID string `xml:",chardata"` // check for id
}

type FlagCarried struct {
	XMLName xml.Name `xml:"flag-carried"`
	BaseFilter
	FlagID string `xml:",chardata"` // check for id
}

type FlagDropped struct {
	XMLName xml.Name `xml:"flag-dropped"`
	BaseFilter
	FlagID string `xml:",chardata"` // check for id
}

type FlagReturned struct {
	XMLName xml.Name `xml:"flag-returned"`
	BaseFilter
	FlagID string `xml:",chardata"` // check for id
}

type FlagCaptured struct {
	XMLName xml.Name `xml:"flag-captured"`
	BaseFilter
	FlagID string `xml:",chardata"` // check for id
}

// Spacial

type Void struct {
	XMLName xml.Name `xml:"void"`
	BaseFilter
}

type Blocks struct {
	XMLName xml.Name `xml:"blocks"`
	BaseFilter

	Region regions.RegionRef `xml:"region,attr"` // check for region
	Filter FilterElement     `xml:",any"`
}

// Blocks

type Material struct {
	XMLName xml.Name `xml:"material"`
	BaseFilter
	Value string `xml:",chardata"` // check for material
}

type StructuralLoad struct {
	XMLName xml.Name `xml:"structural-load"`
	BaseFilter
	Value int `xml:",chardata"`
}

// Entity

type Spawn struct {
	XMLName xml.Name `xml:"spawn"`
	BaseFilter
	Reason string `xml:",chardata"` // check for reason
}

type Mob struct {
	XMLName xml.Name `xml:"mob"`
	BaseFilter
	Name string `xml:",chardata"` // check for mob
}

type Entity struct {
	XMLName xml.Name `xml:"entity"`
	BaseFilter
	Name string `xml:",chardata"` // check for entity
}

// Competitor

type CarryingFlag struct {
	XMLName xml.Name `xml:"carrying-flag"`
	BaseFilter
	FlagID string `xml:",chardata"` // check for id
}

type Score struct {
	XMLName xml.Name `xml:"score"`
	BaseFilter
	TeamID *string `xml:"team,attr,omitempty"` //check for team id
	Value  string  `xml:",chardata"`           // check for valid range
}

type Rank struct {
	XMLName xml.Name `xml:"rank"`
	BaseFilter
	TeamID *string `xml:"team,attr,omitempty"` //check for team id
	Value  string  `xml:",chardata"`           // check for valid range
}

// Player

type Participating struct {
	XMLName xml.Name `xml:"participating"`
	BaseFilter
}

type Observing struct {
	XMLName xml.Name `xml:"observing"`
	BaseFilter
}

type KillStreak struct {
	XMLName xml.Name `xml:"kill-streak"`
	BaseFilter
	Min        *string `xml:"min,attr,omitempty"`
	Max        *string `xml:"max,attr,omitempty"`
	Count      *string `xml:"count,attr,omitempty"`
	Repeat     *bool   `xml:"repeat,attr,omitempty"`
	Persistent *bool   `xml:"persistent,attr,omitempty"`
}

type Lives struct {
	XMLName xml.Name `xml:"lives"`
	BaseFilter
	Min   *string `xml:"min,attr,omitempty"`
	Max   *string `xml:"max,attr,omitempty"`
	Count *string `xml:"count,attr,omitempty"`
}

type Crouching struct {
	XMLName xml.Name `xml:"crouching"`
	BaseFilter
}

type Walking struct {
	XMLName xml.Name `xml:"walking"`
	BaseFilter
}

type Sprinting struct {
	XMLName xml.Name `xml:"sprinting"`
	BaseFilter
}

type Grounded struct {
	XMLName xml.Name `xml:"grounded"`
	BaseFilter
}

type Flying struct {
	XMLName xml.Name `xml:"flying"`
	BaseFilter
}

type Alive struct {
	XMLName xml.Name `xml:"alive"`
	BaseFilter
}

type Dead struct {
	XMLName xml.Name `xml:"dead"`
	BaseFilter
}

type CanFly struct {
	XMLName xml.Name `xml:"can-fly"`
	BaseFilter
}

type Team struct {
	XMLName xml.Name `xml:"team"`
	BaseFilter
	TeamID string `xml:",chardata"` // check for team id
}

type Class struct {
	XMLName xml.Name `xml:"class"`
	BaseFilter
	Name string `xml:",chardata"` // check for class
}

type Effect struct {
	XMLName xml.Name `xml:"effect"`
	BaseFilter
	Name string `xml:",chardata"` // check for effect
}

// Item sub-element
type Item struct {
	XMLName  xml.Name `xml:"item"`
	Material string   `xml:"material,attr"`
}

type Carrying struct {
	XMLName xml.Name `xml:"carrying"`
	BaseFilter
	Amount             *string `xml:"amount,attr,omitempty"`
	IgnoreDurability   *bool   `xml:"ignore-durability,attr,omitempty"`
	IgnoreMetadata     *bool   `xml:"ignore-metadata,attr,omitempty"`
	IgnoreName         *bool   `xml:"ignore-name,attr,omitempty"`
	IgnoreEnchantments *bool   `xml:"ignore-enchantments,attr,omitempty"`
	Item               Item    `xml:"item"`
}

type Holding struct {
	XMLName xml.Name `xml:"holding"`
	BaseFilter
	Amount             *string `xml:"amount,attr,omitempty"`
	IgnoreDurability   *bool   `xml:"ignore-durability,attr,omitempty"`
	IgnoreMetadata     *bool   `xml:"ignore-metadata,attr,omitempty"`
	IgnoreName         *bool   `xml:"ignore-name,attr,omitempty"`
	IgnoreEnchantments *bool   `xml:"ignore-enchantments,attr,omitempty"`
	Item               Item    `xml:"item"`
}

type Wearing struct {
	XMLName xml.Name `xml:"wearing"`
	BaseFilter
	Amount             *string `xml:"amount,attr,omitempty"`
	IgnoreDurability   *bool   `xml:"ignore-durability,attr,omitempty"`
	IgnoreMetadata     *bool   `xml:"ignore-metadata,attr,omitempty"`
	IgnoreName         *bool   `xml:"ignore-name,attr,omitempty"`
	IgnoreEnchantments *bool   `xml:"ignore-enchantments,attr,omitempty"`
	Item               Item    `xml:"item"`
}

// Event

type Cause struct {
	XMLName xml.Name `xml:"cause"`
	BaseFilter
	CauseType string `xml:",chardata"` // check for cause
}

type Random struct {
	XMLName xml.Name `xml:"random"`
	BaseFilter
	Value string `xml:",chardata"` // check for decimal or range
}

// Damage

type Relation struct {
	XMLName xml.Name `xml:"relation"`
	BaseFilter
	RelationType string `xml:",chardata"` // check for relation
}
