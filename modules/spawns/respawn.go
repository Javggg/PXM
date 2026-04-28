package spawns

import (
	"encoding/xml"
	"pxm/modules/globals"
)

type Respawns struct {
	XMLName  xml.Name  `xml:"respawns"`
	Respawns []Respawn `xml:"respawn"`
	Delay    *string   `xml:"delay,attr,omitempty"`
	Filter   *string   `xml:"filter,attr,omitempty"`
	Auto     *string   `xml:"auto,attr,omitempty"`
	Blackout *string   `xml:"blackout,attr,omitempty"`
	Spectate *string   `xml:"spectate,attr,omitempty"`
	Bed      *string   `xml:"bed,attr,omitempty"`
	globals.Globals
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
	globals.Globals
}

type Message struct {
	XMLName xml.Name `xml:"message"`
	Value   string   `xml:",innerxml"`
}
