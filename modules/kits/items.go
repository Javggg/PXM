package kits

import (
	"encoding/xml"
	"pxm/modules/filters"
)

type BaseItem struct {
	Slot *string `xml:"slot,attr,omitempty"`
	Name *string `xml:"name,attr,omitempty"`
	Lore *string `xml:"lore,attr,omitempty"`
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	BaseItem
	Material           string                `xml:"material,attr"`
	Color              *string               `xml:"color,attr,omitempty"`
	Projectile         *string               `xml:"projectile,attr,omitempty"`
	Consumable         *string               `xml:"consumable,attr,omitempty"`
	Amount             *string               `xml:"amount,attr,omitempty"`
	Damage             *string               `xml:"damage,attr,omitempty"`
	Unbreakable        *string               `xml:"unbreakable,attr,omitempty"`
	TeamColor          *string               `xml:"team-color,attr,omitempty"`
	Grenade            *string               `xml:"grenade,attr,omitempty"`
	GrenadePower       *string               `xml:"grenade-power,attr,omitempty"`
	GrenadeFire        *string               `xml:"grenade-fire,attr,omitempty"`
	GrenadeDestroy     *string               `xml:"grenade-destroy,attr,omitempty"`
	PreventSharing     *string               `xml:"prevent-sharing,attr,omitempty"`
	ShowEnchantments   *string               `xml:"show-enchantments,attr,omitempty"`
	ShowAttributes     *string               `xml:"show-attributes,attr,omitempty"`
	ShowUnbreakable    *string               `xml:"show-unbreakable,attr,omitempty"`
	ShowCanDestroy     *string               `xml:"show-can-destroy,attr,omitempty"`
	ShowCanPlaceOn     *string               `xml:"show-can-place-on,attr,omitempty"`
	ShowOther          *string               `xml:"show-other,attr,omitempty"`
	Enchantments       []Enchantment         `xml:"enchantment"`
	StoredEnchantments []Enchantment         `xml:"stored-enchantment"`
	Effects            []Effect              `xml:"effect"`
	Attributes         []Attribute           `xml:"attribute"`
	CanDestroy         *CanDestroyOrCanPlace `xml:"can-destroy,omitempty"`
	CanPlaceOn         *CanDestroyOrCanPlace `xml:"can-place-on,omitempty"`
}

type Enchantment struct {
	Level *string `xml:"level,attr,omitempty"`
	Value string  `xml:",innerxml"`
}

type CanDestroyOrCanPlace struct {
	Materials []filters.Material `xml:"material"`
	AllBlocks *any               `xml:"all-blocks,omitempty"`
}

type Head struct {
	XMLName xml.Name `xml:"head"`
	BaseItem
	Name       string  `xml:"name,attr"`
	InlineUUID *string `xml:"uuid,attr,omitempty"`
	InlineSkin *string `xml:"skin,attr,omitempty"`
	UUID       *string `xml:"uuid,omitempty"`
	Skin       *string `xml:"skin,omitempty"`
}

type Banner struct {
	XMLName xml.Name `xml:"banner"`
	BaseItem
	BaseColor *string `xml:"base-color,attr"`
	Layers    []struct {
		XMLName xml.Name `xml:"layer"`
		Pattern string   `xml:"pattern,attr"`
		Color   string   `xml:"color,attr"`
	} `xml:"layer"`
}

type Armor struct {
	BaseItem
	Material     string        `xml:"material,attr"`
	TeamColor    *string       `xml:"team-color,attr,omitempty"`
	Unbreakable  *string       `xml:"unbreakable,attr,omitempty"`
	Locked       *string       `xml:"locked,attr,omitempty"`
	Enchantments []Enchantment `xml:"enchantment"`
	Attributes   []Attribute   `xml:"attribute"`
}

type Book struct {
	XMLName xml.Name `xml:"book"`
	Title   *string  `xml:"title,omitempty"`
	Author  *string  `xml:"author,omitempty"`
	Pages   struct {
		Page *string `xml:"page,omitempty"`
	} `xml:"pages,omitempty"`
}

type Firework struct {
	XMLName    xml.Name `xml:"firework"`
	Power      *string  `xml:"power,attr,omitempty"`
	Explosions []struct {
		Type    *string  `xml:"type,attr,omitempty"`
		Flicker *string  `xml:"flicker,attr,omitempty"`
		Trail   *string  `xml:"trail,attr,omitempty"`
		Colors  []string `xml:"color"`
		Fades   []string `xml:"fade"`
	}
}
