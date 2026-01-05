package itemmods

import (
	"encoding/xml"
	"pxm/modules/kits"
)

type ItemMods struct {
	XMLName xml.Name `xml:"item-mods"`
	Rules   []Rule   `xml:"rule"`
}

type Rule struct {
	Match  Match  `xml:"match"`
	Modify Modify `xml:"modify"`
}

type Match struct {
	Material     *string `xml:"material,omitempty"`
	AllMaterials *string `xml:"all-materials,omitempty"`
	AllBlocks    *string `xml:"all-blocks,omitempty"`
}

type Modify struct {
	Name             *string                    `xml:"name,attr,omitempty"`
	Lore             *string                    `xml:"lore,attr,omitempty"`
	Unbreakable      *string                    `xml:"unbreakable,attr,omitempty"`
	Potion           *string                    `xml:"potion,attr,omitempty"`
	ShowEnchantments *string                    `xml:"show-enchantments,attr,omitempty"`
	ShowAttributes   *string                    `xml:"show-attributes,attr,omitempty"`
	ShowUnbreakable  *string                    `xml:"show-unbreakable,attr,omitempty"`
	ShowCanDestroy   *string                    `xml:"show-can-destroy,attr,omitempty"`
	ShowCanPlaceOn   *string                    `xml:"show-can-place-on,attr,omitempty"`
	ShowOther        *string                    `xml:"show-other,attr,omitempty"`
	Enchantments     []kits.Enchantment         `xml:"enchantment"`
	Effects          []kits.Effect              `xml:"effect"`
	Attributes       []kits.Attribute           `xml:"attribute"`
	CanDestroys      *kits.CanDestroyOrCanPlace `xml:"can-destroy,omitempty"`
	CanPlaceOn       *kits.CanDestroyOrCanPlace `xml:"can-place-on,omitempty"`
}
