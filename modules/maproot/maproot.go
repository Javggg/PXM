package maproot

import (
	"encoding/xml"
	"pxm/modules/actions"
	"pxm/modules/authors"
	"pxm/modules/constants"
	"pxm/modules/filters"
	"pxm/modules/kits"
	"pxm/modules/portals"
	"pxm/modules/regions"
	"pxm/modules/replacements"
	"pxm/modules/spawners"
	"pxm/modules/spawns"
	"pxm/modules/variables"
)

type Map struct {
	XMLName          xml.Name              `xml:"map"`
	Proto            string                `xml:"proto,attr"`
	MinServerVersion *string               `xml:"min-server-version,attr"`
	MaxServerVersion *string               `xml:"max-server-version,attr"`
	Internal         *string               `xml:"internal,attr"`
	Constants        *constants.Constants  `xml:"constants,omitempty"`
	Authors          *authors.Authors      `xml:"authors,omitempty"`
	Contributors     *authors.Contributors `xml:"contributors,omitempty"`
	// Rules
	// Broadcasts
	// Teams
	Kits     *kits.Kits       `xml:"kits,omitempty"`
	Spawns   *spawns.Spawns   `xml:"spawns,omitempty"`
	Respawns *spawns.Respawns `xml:"respawns,omitempty"`
	// Structures
	// WorldBorders
	// Wools
	// Flags
	// ControlPoints
	// Destroyables
	// Cores
	// Modes
	// Payloads
	Portals      *portals.Portals           `xml:"portals,omitempty"`
	Spawners     *spawners.Spawners         `xml:"spawners,omitempty"`
	Regions      *regions.Regions           `xml:"regions,omitempty"`
	Filters      *filters.Filters           `xml:"filters,omitempty"`
	Variables    *variables.Variables       `xml:"variables,omitempty"`
	Actions      *actions.Actions           `xml:"actions,omitempty"`
	Replacements *replacements.Replacements `xml:"replacements,omitempty"`
	// Stats
	// Projectiles
	// Consumables
	// Renewables
	// KillRewards
	// ItemRemove
	// ItemKeep
	// ArmorKeep
	// ToolRepair
}
