package projectiles

import (
	"pxm/modules/filters"
	"pxm/modules/globals"
	"pxm/modules/kits"
)

type ModifyBowProjectile struct {
	Projectile   *string                  `xml:"projectile,omitempty"`
	VelocityMod  *string                  `xml:"velocityMod,omitempty"`
	Potion       *kits.Effect             `xml:"potion,omitempty"`
	PickupFilter *filters.FilterContainer `xml:"pickup-filter,omitempty"`
	globals.Globals
}
