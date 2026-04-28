package repairremovekeep

import (
	"encoding/xml"
	"pxm/modules/globals"
)

type ToolRepair struct {
	XMLName xml.Name `xml:"toolrepair"`
	Tools   []string `xml:"tool"`
	globals.Globals
}

type ItemRemove struct {
	XMLName xml.Name `xml:"itemremove"`
	Items   []string `xml:"item"`
	globals.Globals
}

type ItemKeep struct {
	XMLName xml.Name `xml:"itemkeep"`
	Items   []string `xml:"item"`
	globals.Globals
}

type ArmorKeep struct {
	XMLName xml.Name `xml:"armorkeep"`
	Items   []string `xml:"item"`
	globals.Globals
}
