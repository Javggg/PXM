package flags

import (
	"encoding/xml"
	"pxm/modules/filters"
	"pxm/modules/globals"
	"pxm/modules/kits"
)

type Flags struct {
	XMLName xml.Name `xml:"flags"`
	Flags   []Flag   `xml:"flag"`
	Posts   []Post   `xml:"post"`
	Nets    []Net    `xml:"net"`
	globals.Globals
}

type Flag struct {
	XMLName                 xml.Name                 `xml:"flag"`
	ID                      *string                  `xml:"id,attr,omitempty"`
	Required                *string                  `xml:"required,attr,omitempty"`
	Name                    *string                  `xml:"name,attr,omitempty"`
	Color                   *string                  `xml:"color,attr,omitempty"`
	ShowMessages            *string                  `xml:"show-messages,attr,omitempty"`
	ShowEffects             *string                  `xml:"show-effects,attr,omitempty"`
	ShowInfo                *string                  `xml:"show-info,attr,omitempty"`
	ShowSidebar             *string                  `xml:"show-sidebar,attr,omitempty"`
	Stats                   *string                  `xml:"stats,attr,omitempty"`
	Show                    *string                  `xml:"show,attr,omitempty"`
	Post                    *string                  `xml:"post,attr,omitempty"`
	Owner                   *string                  `xml:"owner,attr,omitempty"`
	Shared                  *string                  `xml:"shared,attr,omitempty"`
	Sequential              *string                  `xml:"sequential,attr,omitempty"`
	CarryMessage            *string                  `xml:"carry-message,attr,omitempty"`
	ShowRespawnOnPickup     *string                  `xml:"show-respawn-on-pickup,attr,omitempty"`
	Points                  *string                  `xml:"points,attr,omitempty"`
	PointsRate              *string                  `xml:"points-rate,attr,omitempty"`
	InlinePickupFilter      *string                  `xml:"pickup-filter,attr,omitempty"`
	InlineDropFilter        *string                  `xml:"drop-filter,attr,omitempty"`
	InlineCaptureFilter     *string                  `xml:"capture-filter,attr,omitempty"`
	InlinePickupKit         *string                  `xml:"pickup-kit,attr,omitempty"`
	InlineDropKit           *string                  `xml:"drop-kit,attr,omitempty"`
	InlineCarryKit          *string                  `xml:"carry-kit,attr,omitempty"`
	DropOnWater             *string                  `xml:"drop-on-water,attr,omitempty"`
	Beam                    *string                  `xml:"beam,attr,omitempty"`
	FlagProximityMetric     *string                  `xml:"flag-proximity-metric,attr,omitempty"`
	FlagProximityHorizontal *string                  `xml:"flag-proximity-horizontal,attr,omitempty"`
	NetProximityMetric      *string                  `xml:"net-proximity-metric,attr,omitempty"`
	NetProximityHorizontal  *string                  `xml:"net-proximity-horizontal,attr,omitempty"`
	Posts                   []Post                   `xml:"post"`
	Nets                    []Net                    `xml:"net"`
	PickupFilter            *filters.FilterContainer `xml:"pickup-filter,omitempty"`
	DropFilter              *filters.FilterContainer `xml:"drop-filter,omitempty"`
	CaptureFilter           *filters.FilterContainer `xml:"capture-filter,omitempty"`
	PickupKit               *kits.Kit                `xml:"pickup-kit,omitempty"`
	DropKit                 *kits.Kit                `xml:"drop-kit,omitempty"`
	CarryKit                *kits.Kit                `xml:"carry-kit,omitempty"`
	globals.Globals
}

type Post struct {
	XMLName            xml.Name                 `xml:"post"`
	ID                 *string                  `xml:"id,attr,omitempty"`
	Name               *string                  `xml:"name,attr,omitempty"`
	Owner              *string                  `xml:"owner,attr,omitempty"`
	Permanent          *string                  `xml:"permanent,attr,omitempty"`
	Sequential         *string                  `xml:"sequential,attr,omitempty"`
	PointsRate         *string                  `xml:"points-rate,attr,omitempty"`
	InlinePickupFilter *string                  `xml:"pickup-filter,attr,omitempty"`
	RecoverTime        *string                  `xml:"recover-time,attr,omitempty"`
	RespawnTime        *string                  `xml:"respawn-time,attr,omitempty"`
	RespawnFilter      *string                  `xml:"respawn-filter,attr,omitempty"`
	RespawnSpeed       *string                  `xml:"respawn-speed,attr,omitempty"`
	Yaw                *string                  `xml:"yaw,attr,omitempty"`
	Fallback           *string                  `xml:"fallback,attr,omitempty"`
	PickupFilter       *filters.FilterContainer `xml:"pickup-filter,omitempty"`
	Value              string
	Posts              []Post
}

type Net struct {
	XMLName         xml.Name `xml:"net"`
	ID              *string  `xml:"id,attr,omitempty"`
	Region          *string  `xml:"region,attr,omitempty"`
	Owner           *string  `xml:"owner,attr,omitempty"`
	Points          *string  `xml:"points,attr,omitempty"`
	Post            *string  `xml:"post,attr,omitempty"`
	Flags           *string  `xml:"flags,attr,omitempty"`
	Flag            *string  `xml:"flag,attr,omitempty"`
	Rescue          *string  `xml:"rescue,attr,omitempty"`
	Sticky          *string  `xml:"sticky,attr,omitempty"`
	CaptureFilter   *string  `xml:"capture-filter,attr,omitempty"`
	DenyMessage     *string  `xml:"deny-message,attr,omitempty"`
	RespawnTogether *string  `xml:"respawn-together,attr,omitempty"`
	RespawnFilter   *string  `xml:"respawn-filter,attr,omitempty"`
	RespawnMessage  *string  `xml:"respawn-message,attr,omitempty"`
	Location        *string  `xml:"location,attr,omitempty"`
}
