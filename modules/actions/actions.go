package actions

import (
	"encoding/xml"
	"pxm/modules/filters"
)

type Actions struct {
	XMLName  xml.Name  `xml:"actions"`
	Actions  []Action  `xml:"action"`
	Triggers []Trigger `xml:"trigger"`
}

type BaseAction struct {
	ID *string `xml:"id,attr,omitempty"`
}

type Action struct {
	BaseAction
	Scope           string  `xml:"scope,attr,omitempty"` // TODO: check
	Filter          *string `xml:"filter,attr,omitempty"`
	UntriggerFilter *string `xml:"untrigger-filter,attr,omitempty"`
	Expose          *string `xml:"expose,attr,omitempty"`
	Items           []byte  `xml:",innerxml"`
}

type Trigger struct {
	XMLName      xml.Name                 `xml:"trigger"`
	InlineAction *string                  `xml:"action,attr,omitempty"`
	InlineFilter *string                  `xml:"filter,attr,omitempty"`
	Filter       *filters.FilterContainer `xml:"filter,omitempty"`
	Scope        string                   `xml:"scope,attr"` // TODO: check
	Observers    *string                  `xml:"observers,attr,omitempty"`
	Action       Action                   `xml:"action,omitempty"`
}

// type SwitchScope struct {
// 	XMLName   xml.Name `xml:"switch-scope"`
// 	Inner     string   `xml:"inner,attr"` // TODO: check
// 	Outer     *string  `xml:"outer,attr,omitempty"`
// 	Observers *string  `xml:"observers,attr,omitempty"`
// }

// type Repeat struct {
// 	XMLName xml.Name `xml:"repeat"`
// 	Times   string   `xml:"times,attr"`
// 	Filter  *string  `xml:"filter,attr,omitempty"`
// }

// type Message struct {
// 	XMLName      xml.Name                            `xml:"message"`
// 	Text         *string                             `xml:"text,attr,omitempty"` // TODO: support as child
// 	ActionBar    *string                             `xml:"actionbar,attr,omitempty"`
// 	Title        *string                             `xml:"title,attr,omitempty"`
// 	Subtitle     *string                             `xml:"subtitle,attr,omitempty"`
// 	FadeIn       *string                             `xml:"fade-in,attr,omitempty"` // TODO: check time
// 	Stay         *string                             `xml:"stay,attr,omitempty"`
// 	FadeOut      *string                             `xml:"fade-out,attr,omitempty"`
// 	Replacements *replacements.ReplacementsContainer `xml:"replacements,omitempty"`
// }

// type Sound struct {
// 	XMLName xml.Name `xml:"sound"`
// 	Preset  *string  `xml:"preset,attr,omitempty"`
// 	Key     *string  `xml:"key,attr,omitempty"`
// 	Volume  *string  `xml:"volume,attr,omitempty"`
// 	Pitch   *string  `xml:"pitch,attr,omitempty"`
// }

// type Set struct {
// 	XMLName xml.Name `xml:"set"`
// 	Var     string   `xml:"var,attr"` // TODO: check
// 	Index   *string  `xml:"index,attr,omitempty"`
// 	Value   string   `xml:"value,attr"`
// }

// type KillEntities struct {
// 	XMLName xml.Name `xml:"kill-entities"`
// 	Filter  string   `xml:"filter,attr"`
// }

// type Fill struct {
// 	XMLName      xml.Name                 `xml:"fill"`
// 	InlineRegion *string                  `xml:"region,attr,omitempty"`
// 	Region       *regions.RegionContainer `xml:"region,omitempty"`
// 	Material     string                   `xml:"material"`
// 	Filter       *string                  `xml:"filter,attr"`
// 	Events       *string                  `xml:"events,attr,omitempty"`
// 	Update       *string                  `xml:"update,attr,omitempty"`
// }

// type PasteStructure struct {
// 	XMLName   xml.Name `xml:"paste-structure"`
// 	X         string   `xml:"x,attr"`
// 	Y         string   `xml:"y,attr"`
// 	Z         string   `xml:"z,attr"`
// 	Structure string   `xml:"structure,attr"`
// 	Update    *string  `xml:"update,attr,omitempty"`
// }

// type Item struct {
// 	// TODO: import from real kit item
// }

// type ReplaceItem struct {
// 	XMLName        xml.Name `xml:"replace-item"`
// 	Find           Item     `xml:"find"`
// 	Replace        Item     `xml:"replace"`
// 	KeepAmount     *string  `xml:"keep-amount,attr,omitempty"`
// 	KeepEnchants   *string  `xml:"keep-enchants,attr,omitempty"`
// 	IgnoreMetadata *string  `xml:"ignore-metadata,attr,omitempty"`
// 	Amount         *string  `xml:"amount,attr,omitempty"`
// }

// type EnchantItem struct {
// 	XMLName        xml.Name `xml:"enchant-item"`
// 	Enchantment    string   `xml:"enchantment,attr"`
// 	Level          string   `xml:"level,attr"`
// 	IgnoreMetadata string   `xml:"ignore-metadata,attr,omitempty"`
// }

// // TODO: type TakePayment struct {}

// type TeamAlias struct {
// 	XMLName xml.Name `xml:"team-alias"`
// 	Team    *string  `xml:"team,attr,omitempty"`
// 	Alias   string   `xml:"alias,attr"`
// }

// type Velocity struct {
// 	XMLName xml.Name `xml:"velocity"`
// 	X       string   `xml:"x,attr"`
// 	Y       string   `xml:"y,attr"`
// 	Z       string   `xml:"z,attr"`
// 	Yaw     *string  `xml:"yaw,attr,omitempty"`
// 	Pitch   *string  `xml:"pitch,attr,omitempty"`
// }

// type Teleport struct {
// 	XMLName xml.Name `xml:"teleport"`
// 	X       string   `xml:"x,attr"`
// 	Y       string   `xml:"y,attr"`
// 	Z       string   `xml:"z,attr"`
// 	Yaw     *string  `xml:"yaw,attr,omitempty"`
// 	Pitch   *string  `xml:"pitch,attr,omitempty"`
// }

// type Weather struct {
// 	XMLName xml.Name `xml:"weather"`
// 	State   string   `xml:"state,attr"` // TODO: check
// }
