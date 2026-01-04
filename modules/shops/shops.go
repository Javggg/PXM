package shops

import (
	"encoding/xml"
	"pxm/modules/filters"
	"pxm/modules/kits"
)

type Shops struct {
	XMLName xml.Name `xml:"shops"`
	Shops   []Shop   `xml:"shop"`
}

type Shop struct {
	ID         string     `xml:"id,attr"`
	Name       *string    `xml:"name,attr,omitempty"`
	Categories []Category `xml:"category"`
}

type Category struct {
	ID           string                   `xml:"id,attr"`
	Name         *string                  `xml:"name,attr,omitempty"`
	Material     *string                  `xml:"material,attr,omitempty"`
	InlineFilter *string                  `xml:"filter,attr,omitempty"`
	Filter       *filters.FilterContainer `xml:"filter,omitempty"`
	Items        []Item                   `xml:"item"`
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
	} `xml:"payment"`
}
