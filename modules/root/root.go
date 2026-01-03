package root

import (
	"encoding/xml"
	"pxm/modules/actions"
	"pxm/modules/authors"
	"pxm/modules/constants"
	"pxm/modules/consumables"
	"pxm/modules/controlpoints"
	"pxm/modules/cores"
	"pxm/modules/destroyables"
	"pxm/modules/filters"
	"pxm/modules/flags"
	"pxm/modules/killrewards"
	"pxm/modules/kits"
	"pxm/modules/modes"
	"pxm/modules/portals"
	"pxm/modules/projectiles"
	"pxm/modules/regions"
	"pxm/modules/replacements"
	"pxm/modules/spawners"
	"pxm/modules/spawns"
	"pxm/modules/structures"
	"pxm/modules/teams"
	"pxm/modules/variables"
	"pxm/modules/wools"
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
	Teams      *teams.Teams           `xml:"teams,omitempty"`
	Kits       *kits.Kits             `xml:"kits,omitempty"`
	Spawns     *spawns.Spawns         `xml:"spawns,omitempty"`
	Respawns   *spawns.Respawns       `xml:"respawns,omitempty"`
	Structures *structures.Structures `xml:"structures,omitempty"`
	// WorldBorders
	Wools         *wools.Wools                 `xml:"wools,omitempty"`
	Flags         *flags.Flags                 `xml:"flags,omitempty"`
	ControlPoints *controlpoints.ControlPoints `xml:"control-points,omitempty"`
	Destroyables  *destroyables.Destroyables   `xml:"destroyables,omitempty"`
	Cores         *cores.Cores                 `xml:"cores,omitempty"`
	Modes         *modes.Modes                 `xml:"modes,omitempty"`
	// Payloads
	Portals      *portals.Portals           `xml:"portals,omitempty"`
	Spawners     *spawners.Spawners         `xml:"spawners,omitempty"`
	Regions      *regions.Regions           `xml:"regions,omitempty"`
	Filters      *filters.Filters           `xml:"filters,omitempty"`
	Variables    *variables.Variables       `xml:"variables,omitempty"`
	Actions      *actions.Actions           `xml:"actions,omitempty"`
	Replacements *replacements.Replacements `xml:"replacements,omitempty"`
	// Stats
	Projectiles         *projectiles.Projectiles         `xml:"projectiles,omitempty"`
	ModifyBowProjectile *projectiles.ModifyBowProjectile `xml:"modifybowprojectile,omitempty"`
	Consumables         *consumables.Consumables         `xml:"consumables,omitempty"`
	// Renewables
	KillRewards *killrewards.KillRewards `xml:"kill-rewards,omitempty"`
	// ItemRemove
	// ItemKeep
	// ArmorKeep
	// ToolRepair
}
