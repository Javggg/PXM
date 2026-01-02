package cores

import (
	"encoding/xml"
	"pxm/modules/regions"
)

type Cores struct {
	XMLName             xml.Name `xml:"cores"`
	Name                *string  `xml:"name,attr,omitempty"`
	Required            *string  `xml:"required,attr,omitempty"`
	Leak                *string  `xml:"leak,attr,omitempty"`
	Team                *string  `xml:"team,attr,omitempty"`
	Modes               *string  `xml:"modes,attr,omitempty"`
	ModeChanges         *string  `xml:"mode-changes,attr,omitempty"`
	ShowProgress        *string  `xml:"show-progress,attr,omitempty"`
	Material            *string  `xml:"material,attr,omitempty"`
	ShowMessages        *string  `xml:"show-messages,attr,omitempty"`
	ShowEffects         *string  `xml:"show-effects,attr,omitempty"`
	ShowInfo            *string  `xml:"show-info,attr,omitempty"`
	ShowSidebar         *string  `xml:"show-sidebar,attr,omitempty"`
	Stats               *string  `xml:"stats,attr,omitempty"`
	Show                *string  `xml:"show,attr,omitempty"`
	ScoreboardFilter    *string  `xml:"scoreboard-filter,attr,omitempty"`
	ProximityMetric     *string  `xml:"proximity-metric,attr,omitempty"`
	ProximityHorizontal *string  `xml:"proximity-horizontal,attr,omitempty"`
	Cores               []Core   `xml:"core"`
}

type Core struct {
	XMLName             xml.Name                 `xml:"core"`
	ID                  *string                  `xml:"id,attr,omitempty"`
	Name                *string                  `xml:"name,attr,omitempty"`
	Required            *string                  `xml:"required,attr,omitempty"`
	InlineRegion        *string                  `xml:"region,attr,omitempty"`
	Leak                *string                  `xml:"leak,attr,omitempty"`
	Team                *string                  `xml:"team,attr,omitempty"`
	Modes               *string                  `xml:"modes,attr,omitempty"`
	ModeChanges         *string                  `xml:"mode-changes,attr,omitempty"`
	ShowProgress        *string                  `xml:"show-progress,attr,omitempty"`
	Material            *string                  `xml:"material,attr,omitempty"`
	ShowMessages        *string                  `xml:"show-messages,attr,omitempty"`
	ShowEffects         *string                  `xml:"show-effects,attr,omitempty"`
	ShowInfo            *string                  `xml:"show-info,attr,omitempty"`
	ShowSidebar         *string                  `xml:"show-sidebar,attr,omitempty"`
	Stats               *string                  `xml:"stats,attr,omitempty"`
	Show                *string                  `xml:"show,attr,omitempty"`
	ScoreboardFilter    *string                  `xml:"scoreboard-filter,attr,omitempty"`
	ProximityMetric     *string                  `xml:"proximity-metric,attr,omitempty"`
	ProximityHorizontal *string                  `xml:"proximity-horizontal,attr,omitempty"`
	Region              *regions.RegionContainer `xml:"region,omitempty"`
}
