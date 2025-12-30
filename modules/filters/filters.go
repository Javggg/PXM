package filters

import (
	"encoding/xml"
	"pxm/modules/regions"
)

type FilterElement any

type Filters struct {
	XMLName xml.Name `xml:"filters"`
	FilterContainer
}

type BaseFilter struct {
	ID *string `xml:"id,attr,omitempty"`
}

type FilterContainer struct {
	FilterRefs []FilterRef `xml:"filter"`

	Always        []Always        `xml:"always"`
	Never         []Never         `xml:"never"`
	Time          []Time          `xml:"time"`
	MatchStarted  []MatchStarted  `xml:"match-started"`
	MatchRunning  []MatchRunning  `xml:"match-running"`
	MatchFinished []MatchFinished `xml:"match-finished"`
	Completed     []Completed     `xml:"completed"`
	Captured      []Captured      `xml:"captured"`

	FlagCarried  []FlagCarried  `xml:"flag-carried"`
	FlagDropped  []FlagDropped  `xml:"flag-dropped"`
	FlagReturned []FlagReturned `xml:"flag-returned"`
	FlagCaptured []FlagCaptured `xml:"flag-captured"`

	Void   []Void   `xml:"void"`
	Blocks []Blocks `xml:"blocks"`

	Material       []Material       `xml:"material"`
	StructuralLoad []StructuralLoad `xml:"structural-load"`

	Spawn  []Spawn  `xml:"spawn"`
	Mob    []Mob    `xml:"mob"`
	Entity []Entity `xml:"entity"`

	CarryingFlag []CarryingFlag `xml:"carrying-flag"`
	Score        []Score        `xml:"score"`
	Rank         []Rank         `xml:"rank"`

	Participating []Participating `xml:"participating"`
	Observing     []Observing     `xml:"observing"`
	KillStreak    []KillStreak    `xml:"kill-streak"`
	Lives         []Lives         `xml:"lives"`
	Crouching     []Crouching     `xml:"crouching"`
	Walking       []Walking       `xml:"walking"`
	Sprinting     []Sprinting     `xml:"sprinting"`
	Grounded      []Grounded      `xml:"grounded"`
	Flying        []Flying        `xml:"flying"`
	Alive         []Alive         `xml:"alive"`
	Dead          []Dead          `xml:"dead"`
	CanFly        []CanFly        `xml:"can-fly"`
	Team          []Team          `xml:"team"`
	Class         []Class         `xml:"class"`
	Effect        []Effect        `xml:"effect"`

	Carrying []Carrying `xml:"carrying"`
	Holding  []Holding  `xml:"holding"`
	Wearing  []Wearing  `xml:"wearing"`

	Cause  []Cause  `xml:"cause"`
	Random []Random `xml:"random"`

	Relation []Relation `xml:"relation"`

	Countdown []Countdown `xml:"countdown"`
	After     []After     `xml:"after"`
	Pulse     []Pulse     `xml:"pulse"`

	Variable []Variable `xml:"variable"`
	Offset   []Offset   `xml:"offset"`
}

type FilterInPlace struct {
	XMLName xml.Name `xml:"filter"`
	BaseFilter
	FilterContainer
}

type FilterRef struct {
	XMLName xml.Name `xml:"filter"`
	RefID   string   `xml:"id,attr"`
}

type AnyFilter struct { // TODO: custom unmarshaller for this crap
	FilterElement
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
	Value string `xml:",chardata"` // TODO: check for valid duration
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

type Completed struct {
	XMLName xml.Name `xml:"completed"`
	BaseFilter
	Any         *bool   `xml:"any,attr,omitempty"`
	Team        *string `xml:"team,attr,omitempty"` // TODO: check
	ObjectiveID string  `xml:",chardata"`           // TODO: check
}

type Captured struct {
	XMLName xml.Name `xml:"completed"`
	BaseFilter
	ObjectiveID string `xml:",chardata"`
}

type FlagCarried struct {
	XMLName xml.Name `xml:"flag-carried"`
	BaseFilter
	FlagID string `xml:",chardata"` // TODO: check for id
}

type FlagDropped struct {
	XMLName xml.Name `xml:"flag-dropped"`
	BaseFilter
	FlagID string `xml:",chardata"` // TODO: check for id
}

type FlagReturned struct {
	XMLName xml.Name `xml:"flag-returned"`
	BaseFilter
	FlagID string `xml:",chardata"` // TODO: check for id
}

type FlagCaptured struct {
	XMLName xml.Name `xml:"flag-captured"`
	BaseFilter
	FlagID string `xml:",chardata"` // TODO: check for id
}

// Spatial

type Void struct {
	XMLName xml.Name `xml:"void"`
	BaseFilter
}

type Blocks struct {
	XMLName xml.Name `xml:"blocks"`
	BaseFilter

	Region regions.RegionRef `xml:"region,attr"`
	FilterContainer
}

// Blocks

type Material struct {
	XMLName xml.Name `xml:"material"`
	BaseFilter
	Value string `xml:",chardata"` // TODO: check for material
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
	Reason string `xml:",chardata"` // TODO: check for reason
}

type Mob struct {
	XMLName xml.Name `xml:"mob"`
	BaseFilter
	Name string `xml:",chardata"` // TODO: check for mob
}

type Entity struct {
	XMLName xml.Name `xml:"entity"`
	BaseFilter
	Name string `xml:",chardata"` // TODO: check for entity
}

// Competitor

type CarryingFlag struct {
	XMLName xml.Name `xml:"carrying-flag"`
	BaseFilter
	FlagID string `xml:",chardata"` // TODO: check for id
}

type Score struct {
	XMLName xml.Name `xml:"score"`
	BaseFilter
	TeamID *string `xml:"team,attr,omitempty"` // TODO: check for team id
	Value  string  `xml:",chardata"`           // TODO: check for valid range
}

type Rank struct {
	XMLName xml.Name `xml:"rank"`
	BaseFilter
	TeamID *string `xml:"team,attr,omitempty"` // TODO: check for team id
	Value  string  `xml:",chardata"`           // TODO: check for valid range
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
	TeamID string `xml:",chardata"` // TODO: check for team id
}

type Class struct {
	XMLName xml.Name `xml:"class"`
	BaseFilter
	Name string `xml:",chardata"` // TODO: check for class
}

type Effect struct {
	XMLName xml.Name `xml:"effect"`
	BaseFilter
	Name string `xml:",chardata"` // TODO: check for effect
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
	CauseType string `xml:",chardata"` // TODO: check for cause
}

type Random struct {
	XMLName xml.Name `xml:"random"`
	BaseFilter
	Value string `xml:",chardata"` // TODO: check for decimal or range
}

// Damage

type Relation struct {
	XMLName xml.Name `xml:"relation"`
	BaseFilter
	RelationType string `xml:",chardata"` // TODO: check for relation
}

// Time

type Countdown struct {
	XMLName xml.Name `xml:"countdown"`
	BaseFilter
	Duration string     `xml:"duration,attr"`
	Message  *string    `xml:"message,attr,omitempty"`
	FilterID *string    `xml:"filter,attr,omitempty"`
	Filter   *AnyFilter `xml:",any"`
}

type After struct {
	XMLName xml.Name `xml:"after"`
	BaseFilter
	Duration string     `xml:"duration,attr"`
	Message  *string    `xml:"message,attr,omitempty"`
	FilterID *string    `xml:"filter,attr,omitempty"`
	Filter   *AnyFilter `xml:",any"`
}

type Pulse struct {
	XMLName xml.Name `xml:"pulse"`
	BaseFilter
	Period   *string    `xml:"period,attr,omitempty"`
	Duration string     `xml:"duration,attr"`
	FilterID *string    `xml:"filter,attr,omitempty"`
	Filter   *AnyFilter `xml:",any"`
}

// Other

// TODO: check for variable
type Variable struct {
	XMLName xml.Name `xml:"variable"`
	BaseFilter
	VarID string `xml:"var,attr"`
	Value string `xml:",chardata"` // TODO: check for decimal or range
}

type Offset struct {
	XMLName xml.Name `xml:"offset"`
	BaseFilter
	Vector string   `xml:"vector,attr"`
	Filter Material `xml:"material"`
}

type Players struct {
	XMLName xml.Name `xml:"players"`
	BaseFilter
	Min          *int       `xml:"min,attr,omitempty"`
	Max          *int       `xml:"max,attr,omitempty"`
	Participants *bool      `xml:"participants,attr,omitempty"`
	Observers    *bool      `xml:"observers,attr,omitempty"`
	FilterID     *string    `xml:"filter,attr,omitempty"`
	Filter       *AnyFilter `xml:",any"`
}

// Modifiers

type Not struct {
	XMLName xml.Name `xml:"not"`
	BaseFilter
	FilterContainer
}

type One struct {
	XMLName xml.Name `xml:"one"`
	BaseFilter
	FilterContainer
}

type All struct {
	XMLName xml.Name `xml:"all"`
	BaseFilter
	FilterContainer
}

type Any struct {
	XMLName xml.Name `xml:"any"`
	BaseFilter
	FilterContainer
}

type Allow struct {
	XMLName xml.Name `xml:"allow"`
	BaseFilter
	FilterContainer
}

type Deny struct {
	XMLName xml.Name `xml:"deny"`
	BaseFilter
	FilterContainer
}

type SameTeam struct {
	XMLName xml.Name `xml:"same-team"`
	BaseFilter
	FilterContainer
}

type Victim struct {
	XMLName xml.Name `xml:"victim"`
	BaseFilter
	FilterContainer
}

type Attacker struct {
	XMLName xml.Name `xml:"attacker"`
	BaseFilter
	FilterContainer
}

type Player struct {
	XMLName xml.Name `xml:"player"`
	BaseFilter
	FilterContainer
}
