package spawns

import "encoding/xml"

type Respawns struct {
	XMLName  xml.Name  `xml:"respawns"`
	Respawns []Respawn `xml:"respawn"`
	Delay    *string   `xml:"delay,attr,omitempty"`
	Filter   *string   `xml:"filter,attr,omitempty"`
	Auto     *string   `xml:"auto,attr,omitempty"`
	Blackout *string   `xml:"blackout,attr,omitempty"`
	Spectate *string   `xml:"spectate,attr,omitempty"`
	Bed      *string   `xml:"bed,attr,omitempty"`
}

type Respawn struct {
	XMLName  xml.Name `xml:"respawn"`
	Delay    *string  `xml:"delay,attr,omitempty"`
	Filter   *string  `xml:"filter,attr,omitempty"`
	Auto     *string  `xml:"auto,attr,omitempty"`
	Blackout *string  `xml:"blackout,attr,omitempty"`
	Spectate *string  `xml:"spectate,attr,omitempty"`
	Bed      *string  `xml:"bed,attr,omitempty"`
	Message  Message  `xml:"message,omitempty"`
}

type Message struct {
	XMLName xml.Name `xml:"message"`
	Value   string   `xml:",innerxml"`
}
