package shops

import (
	"encoding/xml"
	"pxm/modules/globals"
	"pxm/modules/regions"
)

type Shopkeepers struct {
	XMLName     xml.Name     `xml:"shopkeepers"`
	Shopkeepers []Shopkeeper `xml:"shopkeeper"`
	globals.Globals
}

type Shopkeeper struct {
	Name  *string       `xml:"name,attr,omitempty"`
	Mob   *string       `xml:"mob,attr,omitempty"`
	Shop  string        `xml:"shop,attr"`
	Point regions.Point `xml:"point"`
	globals.Globals
}
