package info

import (
	"encoding/xml"
	"pxm/modules/globals"
)

type Info struct {
	Name      string  `xml:"name"`
	Slug      *string `xml:"slug,omitempty"`
	Version   string  `xml:"version"`
	Objective string  `xml:"objective"`
	Created   *string `xml:"created,omitempty"`
	Phase     *string `xml:"phase,omitempty"`
	Edition   *string `xml:"edition,omitempty"`
	Game      *string `xml:"game,omitempty"`
}

type Authors struct {
	XMLName xml.Name      `xml:"authors"`
	Authors []Contributor `xml:"author"`
	globals.Globals
}

type Contributors struct {
	XMLName      xml.Name      `xml:"contributors"`
	Contributors []Contributor `xml:"contributor"`
	globals.Globals
}

type Contributor struct {
	UUID         *string `xml:"uuid,attr,omitempty"`
	Contribution *string `xml:"contribution,attr,omitempty"`
	Value        *string `xml:",innerxml"`
}

type Include struct {
	XMLName xml.Name `xml:"include"`
	ID      string   `xml:"id,attr"`
}

type Variant struct {
	XMLName  xml.Name `xml:"variant"`
	ID       string   `xml:"id,attr"`
	Override *string  `xml:"override,attr,omitempty"`
	World    *string  `xml:"world,attr,omitempty"`
	Name     string   `xml:",innerxml"`
}
