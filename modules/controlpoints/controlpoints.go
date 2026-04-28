package controlpoints

import (
	"encoding/xml"
	"pxm/modules/filters"
	"pxm/modules/globals"
	"pxm/modules/regions"
)

type ControlPoints struct {
	XMLName          xml.Name       `xml:"control-points"`
	Name             *string        `xml:"name,attr,omitempty"`
	Required         *string        `xml:"required,attr,omitempty"`
	CaptureTime      *string        `xml:"capture-time,attr,omitempty"`
	Points           *string        `xml:"points,attr,omitempty"`
	PointsGrowth     *string        `xml:"points-growth,attr,omitempty"`
	OwnerPoints      *string        `xml:"owner-points,attr,omitempty"`
	CaptureRule      *string        `xml:"capture-rule,attr,omitempty"`
	TimeMultiplier   *string        `xml:"time-multiplier,attr,omitempty"`
	ShowProgress     *string        `xml:"show-progress,attr,omitempty"`
	NeutralState     *string        `xml:"neutral-state,attr,omitempty"`
	RecoveryRate     *string        `xml:"recovery-rate,attr,omitempty"`
	DecayRate        *string        `xml:"decay-rate,attr,omitempty"`
	ContestedRate    *string        `xml:"contested-rate,attr,omitempty"`
	OwnedDecayRate   *string        `xml:"owned-decay-rate,attr,omitempty"`
	Permanent        *string        `xml:"permanent,attr,omitempty"`
	InitialOwner     *string        `xml:"initial-owner,attr,omitempty"`
	ShowMessages     *string        `xml:"show-messages,attr,omitempty"`
	ShowEffects      *string        `xml:"show-effects,attr,omitempty"`
	ShowInfo         *string        `xml:"show-info,attr,omitempty"`
	ShowSidebar      *string        `xml:"show-sidebar,attr,omitempty"`
	Stats            *string        `xml:"stats,attr,omitempty"`
	Show             *string        `xml:"show,attr,omitempty"`
	ScoreboardFilter *string        `xml:"scoreboard-filter,attr,omitempty"`
	VisualMaterials  *string        `xml:"visual-materials,attr,omitempty"`
	Capture          *string        `xml:"capture,attr,omitempty"`
	Progress         *string        `xml:"progress,attr,omitempty"`
	Captured         *string        `xml:"captured,attr,omitempty"`
	CaptureFilter    *string        `xml:"capture-filter,attr,omitempty"`
	PlayerFilter     *string        `xml:"player-filter,attr,omitempty"`
	ControlPoints    []ControlPoint `xml:"control-point"`
	globals.Globals
}

type ControlPoint struct {
	XMLName             xml.Name                 `xml:"control-point"`
	ID                  *string                  `xml:"id,attr,omitempty"`
	Name                *string                  `xml:"name,attr,omitempty"`
	Required            *string                  `xml:"required,attr,omitempty"`
	CaptureTime         *string                  `xml:"capture-time,attr,omitempty"`
	Points              *string                  `xml:"points,attr,omitempty"`
	PointsGrowth        *string                  `xml:"points-growth,attr,omitempty"`
	OwnerPoints         *string                  `xml:"owner-points,attr,omitempty"`
	CaptureRule         *string                  `xml:"capture-rule,attr,omitempty"`
	TimeMultiplier      *string                  `xml:"time-multiplier,attr,omitempty"`
	ShowProgress        *string                  `xml:"show-progress,attr,omitempty"`
	NeutralState        *string                  `xml:"neutral-state,attr,omitempty"`
	RecoveryRate        *string                  `xml:"recovery-rate,attr,omitempty"`
	DecayRate           *string                  `xml:"decay-rate,attr,omitempty"`
	ContestedRate       *string                  `xml:"contested-rate,attr,omitempty"`
	OwnedDecayRate      *string                  `xml:"owned-decay-rate,attr,omitempty"`
	Permanent           *string                  `xml:"permanent,attr,omitempty"`
	InitialOwner        *string                  `xml:"initial-owner,attr,omitempty"`
	ShowMessages        *string                  `xml:"show-messages,attr,omitempty"`
	ShowEffects         *string                  `xml:"show-effects,attr,omitempty"`
	ShowInfo            *string                  `xml:"show-info,attr,omitempty"`
	ShowSidebar         *string                  `xml:"show-sidebar,attr,omitempty"`
	Stats               *string                  `xml:"stats,attr,omitempty"`
	Show                *string                  `xml:"show,attr,omitempty"`
	ScoreboardFilter    *string                  `xml:"scoreboard-filter,attr,omitempty"`
	VisualMaterials     *string                  `xml:"visual-materials,attr,omitempty"`
	InlineCapture       *string                  `xml:"capture,attr,omitempty"`
	InlineProgress      *string                  `xml:"progress,attr,omitempty"`
	InlineCaptured      *string                  `xml:"captured,attr,omitempty"`
	InlineCaptureFilter *string                  `xml:"capture-filter,attr,omitempty"`
	InlinePlayerFilter  *string                  `xml:"player-filter,attr,omitempty"`
	Capture             *regions.RegionContainer `xml:"capture,omitempty"`
	Progress            *regions.RegionContainer `xml:"progress,omitempty"`
	Captured            *regions.RegionContainer `xml:"captured,omitempty"`
	CaptureFilter       *filters.FilterContainer `xml:"capture-filter,omitempty"`
	PlayerFilter        *filters.FilterContainer `xml:"player-filter,omitempty"`
	globals.Globals
}
