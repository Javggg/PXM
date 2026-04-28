package shops

import (
	"encoding/xml"
	"pxm/modules/filters"
	"pxm/modules/globals"
	"pxm/modules/kits"
)

type Shops struct {
	XMLName xml.Name `xml:"shops"`
	Shops   []Shop   `xml:"shop"`
	globals.Globals
}

type Shop struct {
	ID         string     `xml:"id,attr"`
	Name       *string    `xml:"name,attr,omitempty"`
	Categories []Category `xml:"category"`
	globals.Globals
}

type Category struct {
	ID           string                   `xml:"id,attr"`
	Name         *string                  `xml:"name,attr,omitempty"`
	Material     *string                  `xml:"material,attr,omitempty"`
	InlineFilter *string                  `xml:"filter,attr,omitempty"`
	Filter       *filters.FilterContainer `xml:"filter,omitempty"`
	Items        []Item                   `xml:"item"`
	globals.Globals
}

type Item struct {
	kits.Item
	Currency     *string                  `xml:"currency,attr,omitempty"`
	Price        *string                  `xml:"price,attr,omitempty"`
	Action       *string                  `xml:"action,attr,omitempty"`
	Color        *string                  `xml:"color,attr,omitempty"`
	InlineFilter *string                  `xml:"filter,attr,omitempty"`
	Filter       *filters.FilterContainer `xml:"filter,omitempty"`
	Payments     []struct {
		Price    *string `xml:"price,attr,omitempty"`
		Currency *string `xml:"currency,attr,omitempty"`
		Item     *Item   `xml:"item,omitempty"`
		globals.Globals
	} `xml:"payment"`
	globals.Globals
}
