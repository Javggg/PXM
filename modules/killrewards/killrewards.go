package killrewards

import (
	"encoding/xml"
	"pxm/modules/actions"
	"pxm/modules/filters"
	"pxm/modules/kits"
)

type KillRewards struct {
	XMLName     xml.Name     `xml:"kill-rewards"`
	KillRewards []KillReward `xml:"kill-reward"`
}

type KillReward struct {
	XMLName            xml.Name                 `xml:"kill-reward"`
	InlineFilter       *string                  `xml:"filter,attr,omitempty"`
	InlineKit          *string                  `xml:"kit,attr,omitempty"`
	InlineAction       *string                  `xml:"action,attr,omitempty"`
	InlineVictimAction *string                  `xml:"victim-action,attr,omitempty"`
	Filter             *filters.FilterContainer `xml:"filter,omitempty"`
	Kit                *kits.Kit                `xml:"kit,omitempty"`
	Action             *actions.Action          `xml:"action,omitempty"`
	VictimAction       *actions.Action          `xml:"victim-action,omitempty"`
	Items              []kits.Item              `xml:"item"`
}
