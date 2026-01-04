package shops

import "encoding/xml"

type Shopkeepers struct {
	XMLName     xml.Name     `xml:"shopkeepers"`
	Shopkeepers []Shopkeeper `xml:"shopkeeper"`
}

type Shopkeeper struct {
	Name  *string `xml:"name,attr,omitempty"`
	Mob   *string `xml:"mob,attr,omitempty"`
	Shop  string  `xml:"shop,attr"`
	Point struct {
		Yaw   *string `xml:"yaw,attr,omitempty"`
		Pitch *string `xml:"pitch,attr,omitempty"`
		Value string  `xml:",chardata"`
	} `xml:"point"`
}
