package root

import (
	"encoding/xml"
	"pxm/modules/actions"
	"pxm/modules/constants"
	"pxm/modules/consumables"
	"pxm/modules/controlpoints"
	"pxm/modules/cores"
	"pxm/modules/damage"
	"pxm/modules/destroyables"
	"pxm/modules/filters"
	"pxm/modules/flags"
	"pxm/modules/info"
	"pxm/modules/killrewards"
	"pxm/modules/kits"
	"pxm/modules/modes"
	"pxm/modules/portals"
	"pxm/modules/projectiles"
	"pxm/modules/regions"
	"pxm/modules/renewables"
	"pxm/modules/repairremovekeep"
	"pxm/modules/replacements"
	"pxm/modules/spawners"
	"pxm/modules/spawns"
	"pxm/modules/structures"
	"pxm/modules/teams"
	"pxm/modules/variables"
	"pxm/modules/wools"
	"pxm/modules/worldborders"
)

type Map struct {
	XMLName          xml.Name `xml:"map"`
	Proto            *string  `xml:"proto,attr"`
	MinServerVersion *string  `xml:"min-server-version,attr"`
	MaxServerVersion *string  `xml:"max-server-version,attr"`
	Internal         *string  `xml:"internal,attr"`
	info.Info
	Gamemodes    []string             `xml:"gamemode"`
	Includes     []info.Include       `xml:"include"`
	Authors      *info.Authors        `xml:"authors,omitempty"`
	Contributors *info.Contributors   `xml:"contributors,omitempty"`
	Constants    *constants.Constants `xml:"constants,omitempty"`
	// Rules
	// Broadcasts
	Teams         *teams.Teams                 `xml:"teams,omitempty"`
	Players       *teams.Players               `xml:"players,omitempty"`
	Kits          *kits.Kits                   `xml:"kits,omitempty"`
	Spawns        *spawns.Spawns               `xml:"spawns,omitempty"`
	Respawns      *spawns.Respawns             `xml:"respawns,omitempty"`
	Structures    *structures.Structures       `xml:"structures,omitempty"`
	WorldBorders  *worldborders.WorldBorders   `xml:"world-borders,omitempty"`
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
	Renewables          *renewables.Renewables           `xml:"renewables,omitempty"`
	KillRewards         *killrewards.KillRewards         `xml:"kill-rewards,omitempty"`
	ItemRemove          *repairremovekeep.ItemRemove     `xml:"itemremove,omitempty"`
	ItemKeep            *repairremovekeep.ItemKeep       `xml:"itemkeep,omitempty"`
	ArmorKeep           *repairremovekeep.ArmorKeep      `xml:"armorkeep,omitempty"`
	ToolRepair          *repairremovekeep.ToolRepair     `xml:"toolrepair,omitempty"`
	Damage              *damage.Damage                   `xml:"damage,omitempty"`
	DisableDamage       *damage.DisableDamage            `xml:"disabledamage,omitempty"`
	FriendlyFire        *damage.FriendlyFire             `xml:"friendlyfire,omitempty"`
	FriendlyFireRefund  *damage.FriendlyFireRefund       `xml:"friendlyfirerefund,omitempty"`
	Difficulty          *damage.Difficulty               `xml:"difficulty,omitempty"`
	Hunger              *damage.Hunger                   `xml:"hunger,omitempty"`
}
