package repairremovekeep

import "encoding/xml"

type ToolRepair struct {
	XMLName xml.Name `xml:"toolrepair"`
	Tools   []string `xml:"tool"`
}

type ItemRemove struct {
	XMLName xml.Name `xml:"itemremove"`
	Items   []string `xml:"item"`
}

type ItemKeep struct {
	XMLName xml.Name `xml:"itemkeep"`
	Items   []string `xml:"item"`
}

type ArmorKeep struct {
	XMLName xml.Name `xml:"armorkeep"`
	Items   []string `xml:"item"`
}
