package modules

import "encoding/xml"

type RegionElement interface{}

type Regions struct {
	XMLName xml.Name        `xml:"regions"`
	Items   []RegionElement `xml:",any"`
}

type Cuboid struct {
	XMLName xml.Name `xml:"cuboid"`
	ID      string   `xml:"id,attr,omitempty"`
	Min     string   `xml:"min,attr,omitempty"`
	Max     string   `xml:"max,attr,omitempty"`
	Size    string   `xml:"size,attr,omitempty"`
}

type Cylinder struct {
	XMLName xml.Name `xml:"cylinder"`
	ID      string   `xml:"id,attr,omitempty"`
	Base    string   `xml:"base,attr"`
	Radius  string   `xml:"radius,attr"`
	Height  string   `xml:"height,attr"`
}

type Block struct {
	XMLName xml.Name `xml:"block"`
	ID      string   `xml:"id,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

type Sphere struct {
	XMLName xml.Name `xml:"sphere"`
	ID      string   `xml:"id,attr,omitempty"`
	Origin  string   `xml:"origin,attr"`
	Radius  string   `xml:"radius,attr"`
}

type Point struct {
	XMLName xml.Name `xml:"point"`
	ID      string   `xml:"id,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

type Rectangle struct {
	XMLName xml.Name `xml:"rectangle"`
	ID      string   `xml:"id,attr,omitempty"`
	Min     string   `xml:"min,attr"`
	Max     string   `xml:"max,attr"`
}

type Circle struct {
	XMLName xml.Name `xml:"circle"`
	ID      string   `xml:"id,attr,omitempty"`
	Center  string   `xml:"center,attr"`
	Radius  string   `xml:"radius,attr"`
}

type Half struct {
	XMLName xml.Name `xml:"half"`
	Normal  string   `xml:"normal,attr"`
	Origin  string   `xml:"origin,attr"`
	ID      string   `xml:"id,attr,omitempty"`
}

type Below struct {
	XMLName xml.Name `xml:"below"`
	X       string   `xml:"x,attr,omitempty"`
	Y       string   `xml:"y,attr,omitempty"`
	Z       string   `xml:"z,attr,omitempty"`
	ID      string   `xml:"id,attr,omitempty"`
}

type Above struct {
	XMLName xml.Name `xml:"above"`
	X       string   `xml:"x,attr,omitempty"`
	Y       string   `xml:"y,attr,omitempty"`
	Z       string   `xml:"z,attr,omitempty"`
	ID      string   `xml:"id,attr,omitempty"`
}

type Empty struct {
	XMLName xml.Name `xml:"empty"`
}

type Nowhere struct {
	XMLName xml.Name `xml:"nowhere"`
}

type Everywhere struct {
	XMLName xml.Name `xml:"everywhere"`
}

// Modifiers

type Negative struct {
	XMLName  xml.Name        `xml:"negative"`
	Children []RegionElement `xml:",any"`
}

type Union struct {
	XMLName  xml.Name        `xml:"union"`
	Children []RegionElement `xml:",any"`
}

type Complement struct {
	XMLName  xml.Name        `xml:"complement"`
	Children []RegionElement `xml:",any"`
}

type Intersect struct {
	XMLName  xml.Name        `xml:"intersect"`
	Children []RegionElement `xml:",any"`
}

type Translate struct {
	XMLName  xml.Name        `xml:"translate"`
	Offset   string          `xml:"offset,attr"`
	Children []RegionElement `xml:",any"`
}

type Mirror struct {
	XMLName  xml.Name        `xml:"mirror"`
	Normal   string          `xml:"normal,attr"`
	Origin   string          `xml:"origin,attr"`
	Children []RegionElement `xml:",any"`
}
