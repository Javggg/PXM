package regions

import (
	"encoding/xml"
	"pxm/modules/globals"
)

type Regions struct {
	XMLName xml.Name `xml:"regions"`
	RegionContainer
	Applies []Apply `xml:"apply"`
}

type BaseRegion struct {
	ID *string `xml:"id,attr,omitempty"`
}

type RegionContainer struct {
	RegionRefs []RegionRef `xml:"region"`

	Unions      []Union      `xml:"union"`
	Negatives   []Negative   `xml:"negative"`
	Complements []Complement `xml:"complement"`
	Intersects  []Intersect  `xml:"intersect"`
	Translates  []Translate  `xml:"translate"`
	Mirrors     []Mirror     `xml:"mirror"`

	Cuboids    []Cuboid    `xml:"cuboid"`
	Cylinders  []Cylinder  `xml:"cylinder"`
	Spheres    []Sphere    `xml:"sphere"`
	Points     []Point     `xml:"point"`
	Rectangles []Rectangle `xml:"rectangle"`
	Circles    []Circle    `xml:"circle"`
	Blocks     []Block     `xml:"block"`

	Halves []Half  `xml:"half"`
	Below  []Below `xml:"below"`
	Above  []Above `xml:"above"`

	Empty      []Empty      `xml:"empty"`
	Nowhere    []Nowhere    `xml:"nowhere"`
	Everywhere []Everywhere `xml:"everywhere"`

	globals.Globals
}

type RegionRef struct {
	XMLName xml.Name `xml:"region"`
	RefID   string   `xml:"id,attr"`
}

type Cuboid struct {
	XMLName xml.Name `xml:"cuboid"`
	BaseRegion
	Min  string  `xml:"min,attr,omitempty"`
	Max  string  `xml:"max,attr,omitempty"`
	Size *string `xml:"size,attr,omitempty"`
}

type Cylinder struct {
	XMLName xml.Name `xml:"cylinder"`
	BaseRegion
	Base   string `xml:"base,attr"`
	Radius string `xml:"radius,attr"`
	Height string `xml:"height,attr"`
}

type Block struct {
	XMLName xml.Name `xml:"block"`
	BaseRegion
	Value string `xml:",innerxml"`
}

type Sphere struct {
	XMLName xml.Name `xml:"sphere"`
	BaseRegion
	Origin string `xml:"origin,attr"`
	Radius string `xml:"radius,attr"`
}

type Point struct {
	XMLName xml.Name `xml:"point"`
	BaseRegion
	Yaw   *string `xml:"yaw,attr,omitempty"`
	Pitch *string `xml:"pitch,attr,omitempty"`
	Value string  `xml:",innerxml"`
}

type Rectangle struct {
	XMLName xml.Name `xml:"rectangle"`
	BaseRegion
	Min string `xml:"min,attr"`
	Max string `xml:"max,attr"`
}

type Circle struct {
	XMLName xml.Name `xml:"circle"`
	BaseRegion
	Center string `xml:"center,attr"`
	Radius string `xml:"radius,attr"`
}

type Half struct {
	XMLName xml.Name `xml:"half"`
	BaseRegion
	Normal string `xml:"normal,attr"`
	Origin string `xml:"origin,attr"`
}

type Below struct {
	XMLName xml.Name `xml:"below"`
	BaseRegion
	X string `xml:"x,attr,omitempty"`
	Y string `xml:"y,attr,omitempty"`
	Z string `xml:"z,attr,omitempty"`
}

type Above struct {
	XMLName xml.Name `xml:"above"`
	BaseRegion
	X string `xml:"x,attr,omitempty"`
	Y string `xml:"y,attr,omitempty"`
	Z string `xml:"z,attr,omitempty"`
}

type Empty struct {
	XMLName xml.Name `xml:"empty"`
	BaseRegion
}

type Nowhere struct {
	XMLName xml.Name `xml:"nowhere"`
	BaseRegion
}

type Everywhere struct {
	XMLName xml.Name `xml:"everywhere"`
	BaseRegion
}

// Modifiers

type Negative struct {
	XMLName xml.Name `xml:"negative"`
	BaseRegion
	RegionContainer
}

type Union struct {
	XMLName xml.Name `xml:"union"`
	BaseRegion
	RegionContainer
}

type Complement struct {
	XMLName xml.Name `xml:"complement"`
	BaseRegion
	RegionContainer
}

type Intersect struct {
	XMLName xml.Name `xml:"intersect"`
	BaseRegion
	RegionContainer
}

type Translate struct {
	XMLName xml.Name `xml:"translate"`
	BaseRegion
	Offset string  `xml:"offset,attr"`
	Region *string `xml:"region,attr"`
	RegionContainer
}

type Mirror struct {
	XMLName xml.Name `xml:"mirror"`
	BaseRegion
	Normal string  `xml:"normal,attr"`
	Origin string  `xml:"origin,attr"`
	Region *string `xml:"region,attr"`
	RegionContainer
}
