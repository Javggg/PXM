package authors

import "encoding/xml"

type Authors struct {
	XMLName xml.Name `xml:"authors"`
	Authors []Author `xml:"author"`
}

type Author struct {
	XMLName      xml.Name `xml:"author"`
	UUID         string   `xml:"uuid,attr"`
	Contribution *string  `xml:"contribution,attr"`
}

type Contributors struct {
	XMLName      xml.Name      `xml:"contributors"`
	Contributors []Contributor `xml:"contributor"`
}

type Contributor struct {
	XMLName      xml.Name `xml:"contributor"`
	UUID         string   `xml:"uuid,attr"`
	Contribution *string  `xml:"contribution,attr"`
}
