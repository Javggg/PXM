package destroyables

import (
	"encoding/xml"
	"pxm/modules/regions"
)

type Destroyables struct {
	XMLName             xml.Name      `xml:"destroyables"`
	Name                *string       `xml:"name,attr,omitempty"`
	Required            *string       `xml:"required,attr,omitempty"`
	Materials           *string       `xml:"materials,attr,omitempty"`
	Owner               *string       `xml:"owner,attr,omitempty"`
	Completion          *string       `xml:"completion,attr,omitempty"`
	Modes               *string       `xml:"modes,attr,omitempty"`
	ModeChanges         *string       `xml:"mode-changes,attr,omitempty"`
	ShowProgress        *string       `xml:"show-progress,attr,omitempty"`
	Repairable          *string       `xml:"repairable,attr,omitempty"`
	Sparks              *string       `xml:"sparks,attr,omitempty"`
	ShowMessages        *string       `xml:"show-messages,attr,omitempty"`
	ShowEffects         *string       `xml:"show-effects,attr,omitempty"`
	ShowInfo            *string       `xml:"show-info,attr,omitempty"`
	ShowSidebar         *string       `xml:"show-sidebar,attr,omitempty"`
	Stats               *string       `xml:"stats,attr,omitempty"`
	Show                *string       `xml:"show,attr,omitempty"`
	ScoreboardFilter    *string       `xml:"scoreboard-filter,attr,omitempty"`
	ProximityMetric     *string       `xml:"proximity-metric,attr,omitempty"`
	ProximityHorizontal *string       `xml:"proximity-horizontal,attr,omitempty"`
	Destroyables        []Destroyable `xml:"destroyable"`
}

type Destroyable struct {
	XMLName             xml.Name                 `xml:"destroyable"`
	ID                  *string                  `xml:"id,attr,omitempty"`
	Name                *string                  `xml:"name,attr,omitempty"`
	Required            *string                  `xml:"required,attr,omitempty"`
	InlineRegion        *string                  `xml:"region,attr,omitempty"`
	Materials           *string                  `xml:"materials,attr,omitempty"`
	Owner               *string                  `xml:"owner,attr,omitempty"`
	Completion          *string                  `xml:"completion,attr,omitempty"`
	Modes               *string                  `xml:"modes,attr,omitempty"`
	ModeChanges         *string                  `xml:"mode-changes,attr,omitempty"`
	ShowProgress        *string                  `xml:"show-progress,attr,omitempty"`
	Repairable          *string                  `xml:"repairable,attr,omitempty"`
	Sparks              *string                  `xml:"sparks,attr,omitempty"`
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
