package kits

import (
	"encoding/xml"
	"pxm/modules/actions"
)

type Kits struct {
	Kits []Kit          `xml:"kit"`
	Give []GiveTakeLend `xml:"give"`
	Take []GiveTakeLend `xml:"take"`
	Lend []GiveTakeLend `xml:"lend"`
}

type Kit struct {
	ID                 *string             `xml:"id,attr,omitempty"`
	Filter             *string             `xml:"filter,attr,omitempty"`
	Parents            *string             `xml:"parents,attr,omitempty"`
	Force              *string             `xml:"force,attr,omitempty"`
	OverflowWarning    *string             `xml:"overflow-warning,attr,omitempty"`
	RepairTools        *string             `xml:"repair-tools,attr,omitempty"`
	DeductTools        *string             `xml:"deduct-tools,attr,omitempty"`
	DeductItems        *string             `xml:"deduct-items,attr,omitempty"`
	DropOverflow       *string             `xml:"drop-overflow,attr,omitempty"`
	PotionParticles    *string             `xml:"potion-particles,attr,omitempty"`
	ResetEnderPearls   *string             `xml:"reset-ender-pearls,attr,omitempty"`
	Clear              *Clear              `xml:"clear,omitempty"`
	Items              []Item              `xml:"item"`
	Helmet             *Armor              `xml:"helmet,omitempty"`
	Chestplate         *Armor              `xml:"chestplate,omitempty"`
	Leggings           *Armor              `xml:"leggings,omitempty"`
	Boots              *Armor              `xml:"boots,omitempty"`
	Effects            []Effect            `xml:"effect"`
	Potions            []Effect            `xml:"potion"`
	Health             *HealthOrFoodLevel  `xml:"health,omitempty"`
	MaxHealth          *HealthOrFoodLevel  `xml:"max-health,omitempty"`
	Saturation         *HealthOrFoodLevel  `xml:"saturation,omitempty"`
	FoodLevel          *HealthOrFoodLevel  `xml:"foodlevel,omitempty"`
	DoubleJump         *DoubleJump         `xml:"double-jump,omitempty"`
	Fly                *Fly                `xml:"fly,omitempty"`
	Shield             *Shield             `xml:"shield,omitempty"`
	WalkSpeed          *WalkSpeed          `xml:"walk-speed,omitempty"`
	KnockbackReduction *KnockbackReduction `xml:"knockback-reduction,omitempty"`
	TeamSwitch         *TeamSwitch         `xml:"team-switch,omitempty"`
	GameMode           *GameMode           `xml:"game-mode,omitempty"`
	Attribute          []Attribute         `xml:"attribute"`
	Kits               []Kit               `xml:"kit"`
	Action             []*actions.Action   `xml:"action"`
}

type GiveTakeLend struct {
	Kit    string `xml:"kit,attr"`
	Filter string `xml:"filter,attr"`
}

type Clear struct {
	XMLName xml.Name `xml:"clear"`
	Items   *string  `xml:"items,attr,omitempty"`
	Armor   *string  `xml:"armor,attr,omitempty"`
	Effects *string  `xml:"effects,attr,omitempty"`
}

type GameMode struct {
	XMLName xml.Name `xml:"game-mode"`
	Value   string   `xml:",innerxml"`
}

type HealthOrFoodLevel struct {
	Value string `xml:",innerxml"`
}

type TeamSwitch struct {
	XMLName   xml.Name `xml:"team-switch"`
	Team      string   `xml:"team,attr"`
	ShowTitle *string  `xml:"show-title,attr,omitempty"`
}

type Effect struct {
	Duration  *string `xml:"duration,attr,omitempty"`
	Amplifier *string `xml:"amplifier,attr,omitempty"`
	Ambient   *string `xml:"ambient,attr,omitempty"`
	Value     string  `xml:",innerxml"`
}

type Attribute struct {
	XMLName   xml.Name `xml:"attribute"`
	Operation *string  `xml:"operation,attr,omitempty"`
	Amount    string   `xml:"amount,attr"`
	Slot      *string  `xml:"slot,attr,omitempty"`
	Value     string   `xml:",innerxml"`
}

type WalkSpeed struct {
	XMLName xml.Name `xml:"walk-speed"`
	Value   string   `xml:",innerxml"`
}

type KnockbackReduction struct {
	XMLName xml.Name `xml:"knockback-reduction"`
	Value   string   `xml:",innerxml"`
}

type Shield struct {
	XMLName xml.Name `xml:"shield"`
	Health  *string  `xml:"health,attr,omitempty"`
	Delay   *string  `xml:"delay,attr,omitempty"`
}

type DoubleJump struct {
	XMLName               xml.Name `xml:"double-jump"`
	Enabled               *string  `xml:"enabled,attr,omitempty"`
	Power                 *string  `xml:"power,attr,omitempty"`
	RechargeTime          *string  `xml:"recharge-time,attr,omitempty"`
	RechargeBeforeLanding *string  `xml:"recharge-before-landing,attr,omitempty"`
}

type Fly struct {
	XMLName  xml.Name `xml:"fly"`
	CanFly   *string  `xml:"can-fly,attr,omitempty"`
	Flying   *string  `xml:"flying,attr,omitempty"`
	FlySpeed *string  `xml:"fly-speed,attr,omitempty"`
}
