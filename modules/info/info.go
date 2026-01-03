package info

import "encoding/xml"

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
}

type Contributors struct {
	XMLName      xml.Name      `xml:"contributors"`
	Contributors []Contributor `xml:"contributor"`
}

type Contributor struct {
	UUID         string  `xml:"uuid,attr"`
	Contribution *string `xml:"contribution,attr,omitempty"`
}

type Include struct {
	XMLName xml.Name `xml:"include"`
	ID      string   `xml:"id,attr"`
}
