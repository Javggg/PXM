package root

func (m *Map) Merge(other Map) {
	// Authors
	if m.Authors != nil && other.Authors != nil {
		m.Authors.Merge(*other.Authors)
	} else if m.Authors == nil && other.Authors != nil {
		copy := *other.Authors
		m.Authors = &copy
	}

	// Contributors
	if m.Contributors != nil && other.Contributors != nil {
		m.Contributors.Merge(*other.Contributors)
	} else if m.Contributors == nil && other.Contributors != nil {
		copy := *other.Contributors
		m.Contributors = &copy
	}

	// Gamemodes
	if len(other.Gamemodes) > 0 {
		m.Gamemodes = append(m.Gamemodes, other.Gamemodes...)
	}

	// Variants
	if len(other.Variants) > 0 {
		m.Variants = append(m.Variants, other.Variants...)
	}

	// Includes
	if len(other.Includes) > 0 {
		m.Includes = append(m.Includes, other.Includes...)
	}

	// Constants
	if m.Constants != nil && other.Constants != nil {
		m.Constants.Merge(*other.Constants)
	} else if m.Constants == nil && other.Constants != nil {
		copy := *other.Constants
		m.Constants = &copy
	}

	// If
	if len(other.If) > 0 {
		m.If = append(m.If, other.If...)
	}

	// Unless
	if len(other.Unless) > 0 {
		m.Unless = append(m.Unless, other.Unless...)
	}

	// Teams
	if m.Teams != nil && other.Teams != nil {
		m.Teams.Merge(*other.Teams)
	} else if m.Teams == nil && other.Teams != nil {
		copy := *other.Teams
		m.Teams = &copy
	}

	// Kits
	if m.Kits != nil && other.Kits != nil {
		m.Kits.Merge(*other.Kits)
	} else if m.Kits == nil && other.Kits != nil {
		copy := *other.Kits
		m.Kits = &copy
	}

	// Spawns
	if m.Spawns != nil && other.Spawns != nil {
		m.Spawns.Merge(*other.Spawns)
	} else if m.Spawns == nil && other.Spawns != nil {
		copy := *other.Spawns
		m.Spawns = &copy
	}

	// Respawns
	if m.Respawns != nil && other.Respawns != nil {
		m.Respawns.Merge(*other.Respawns)
	} else if m.Respawns == nil && other.Respawns != nil {
		copy := *other.Respawns
		m.Respawns = &copy
	}

	// Structures
	if m.Structures != nil && other.Structures != nil {
		m.Structures.Merge(*other.Structures)
	} else if m.Structures == nil && other.Structures != nil {
		copy := *other.Structures
		m.Structures = &copy
	}

	// Lootables
	if m.Lootables != nil && other.Lootables != nil {
		m.Lootables.Merge(*other.Lootables)
	} else if m.Lootables == nil && other.Lootables != nil {
		copy := *other.Lootables
		m.Lootables = &copy
	}

	// WorldBorders
	if m.WorldBorders != nil && other.WorldBorders != nil {
		m.WorldBorders.Merge(*other.WorldBorders)
	} else if m.WorldBorders == nil && other.WorldBorders != nil {
		copy := *other.WorldBorders
		m.WorldBorders = &copy
	}

	// Score
	if m.Score != nil && other.Score != nil {
		m.Score.Merge(*other.Score)
	} else if m.Score == nil && other.Score != nil {
		copy := *other.Score
		m.Score = &copy
	}

	// Wools
	if m.Wools != nil && other.Wools != nil {
		m.Wools.Merge(*other.Wools)
	} else if m.Wools == nil && other.Wools != nil {
		copy := *other.Wools
		m.Wools = &copy
	}

	// Flags
	if m.Flags != nil && other.Flags != nil {
		m.Flags.Merge(*other.Flags)
	} else if m.Flags == nil && other.Flags != nil {
		copy := *other.Flags
		m.Flags = &copy
	}

	// ControlPoints
	if m.ControlPoints != nil && other.ControlPoints != nil {
		m.ControlPoints.Merge(*other.ControlPoints)
	} else if m.ControlPoints == nil && other.ControlPoints != nil {
		copy := *other.ControlPoints
		m.ControlPoints = &copy
	}

	// Destroyables
	if m.Destroyables != nil && other.Destroyables != nil {
		m.Destroyables.Merge(*other.Destroyables)
	} else if m.Destroyables == nil && other.Destroyables != nil {
		copy := *other.Destroyables
		m.Destroyables = &copy
	}

	// Cores
	if m.Cores != nil && other.Cores != nil {
		m.Cores.Merge(*other.Cores)
	} else if m.Cores == nil && other.Cores != nil {
		copy := *other.Cores
		m.Cores = &copy
	}

	// Modes
	if m.Modes != nil && other.Modes != nil {
		m.Modes.Merge(*other.Modes)
	} else if m.Modes == nil && other.Modes != nil {
		copy := *other.Modes
		m.Modes = &copy
	}

	// Regions
	if m.Regions != nil && other.Regions != nil {
		m.Regions.Merge(*other.Regions)
	} else if m.Regions == nil && other.Regions != nil {
		copy := *other.Regions
		m.Regions = &copy
	}

	// Portals
	if m.Portals != nil && other.Portals != nil {
		m.Portals.Merge(*other.Portals)
	} else if m.Portals == nil && other.Portals != nil {
		copy := *other.Portals
		m.Portals = &copy
	}

	// Spawners
	if m.Spawners != nil && other.Spawners != nil {
		m.Spawners.Merge(*other.Spawners)
	} else if m.Spawners == nil && other.Spawners != nil {
		copy := *other.Spawners
		m.Spawners = &copy
	}

	// Filters
	if m.Filters != nil && other.Filters != nil {
		m.Filters.Merge(*other.Filters)
	} else if m.Filters == nil && other.Filters != nil {
		copy := *other.Filters
		m.Filters = &copy
	}

	// Variables
	if m.Variables != nil && other.Variables != nil {
		m.Variables.Merge(*other.Variables)
	} else if m.Variables == nil && other.Variables != nil {
		copy := *other.Variables
		m.Variables = &copy
	}

	// Actions
	if m.Actions != nil && other.Actions != nil {
		m.Actions.Merge(*other.Actions)
	} else if m.Actions == nil && other.Actions != nil {
		copy := *other.Actions
		m.Actions = &copy
	}

	// Replacements
	if m.Replacements != nil && other.Replacements != nil {
		m.Replacements.Merge(*other.Replacements)
	} else if m.Replacements == nil && other.Replacements != nil {
		copy := *other.Replacements
		m.Replacements = &copy
	}

	// Shops
	if m.Shops != nil && other.Shops != nil {
		m.Shops.Merge(*other.Shops)
	} else if m.Shops == nil && other.Shops != nil {
		copy := *other.Shops
		m.Shops = &copy
	}

	// ShopKeepers
	if m.ShopKeepers != nil && other.ShopKeepers != nil {
		m.ShopKeepers.Merge(*other.ShopKeepers)
	} else if m.ShopKeepers == nil && other.ShopKeepers != nil {
		copy := *other.ShopKeepers
		m.ShopKeepers = &copy
	}

	// Projectiles
	if m.Projectiles != nil && other.Projectiles != nil {
		m.Projectiles.Merge(*other.Projectiles)
	} else if m.Projectiles == nil && other.Projectiles != nil {
		copy := *other.Projectiles
		m.Projectiles = &copy
	}

	// Consumables
	if m.Consumables != nil && other.Consumables != nil {
		m.Consumables.Merge(*other.Consumables)
	} else if m.Consumables == nil && other.Consumables != nil {
		copy := *other.Consumables
		m.Consumables = &copy
	}

	// Renewables
	if m.Renewables != nil && other.Renewables != nil {
		m.Renewables.Merge(*other.Renewables)
	} else if m.Renewables == nil && other.Renewables != nil {
		copy := *other.Renewables
		m.Renewables = &copy
	}

	// FallingBlocks
	if m.FallingBlocks != nil && other.FallingBlocks != nil {
		m.FallingBlocks.Merge(*other.FallingBlocks)
	} else if m.FallingBlocks == nil && other.FallingBlocks != nil {
		copy := *other.FallingBlocks
		m.FallingBlocks = &copy
	}

	// ItemMods
	if m.ItemMods != nil && other.ItemMods != nil {
		m.ItemMods.Merge(*other.ItemMods)
	} else if m.ItemMods == nil && other.ItemMods != nil {
		copy := *other.ItemMods
		m.ItemMods = &copy
	}

	// KillRewards
	if m.KillRewards != nil && other.KillRewards != nil {
		m.KillRewards.Merge(*other.KillRewards)
	} else if m.KillRewards == nil && other.KillRewards != nil {
		copy := *other.KillRewards
		m.KillRewards = &copy
	}

	// ItemRemove
	if m.ItemRemove != nil && other.ItemRemove != nil {
		m.ItemRemove.Merge(*other.ItemRemove)
	} else if m.ItemRemove == nil && other.ItemRemove != nil {
		copy := *other.ItemRemove
		m.ItemRemove = &copy
	}

	// ItemKeep
	if m.ItemKeep != nil && other.ItemKeep != nil {
		m.ItemKeep.Merge(*other.ItemKeep)
	} else if m.ItemKeep == nil && other.ItemKeep != nil {
		copy := *other.ItemKeep
		m.ItemKeep = &copy
	}

	// ArmorKeep
	if m.ArmorKeep != nil && other.ArmorKeep != nil {
		m.ArmorKeep.Merge(*other.ArmorKeep)
	} else if m.ArmorKeep == nil && other.ArmorKeep != nil {
		copy := *other.ArmorKeep
		m.ArmorKeep = &copy
	}

	// ToolRepair
	if m.ToolRepair != nil && other.ToolRepair != nil {
		m.ToolRepair.Merge(*other.ToolRepair)
	} else if m.ToolRepair == nil && other.ToolRepair != nil {
		copy := *other.ToolRepair
		m.ToolRepair = &copy
	}

	// Compass
	if m.Compass != nil && other.Compass != nil {
		m.Compass.Merge(*other.Compass)
	} else if m.Compass == nil && other.Compass != nil {
		copy := *other.Compass
		m.Compass = &copy
	}

	// DisableDamage
	if m.DisableDamage != nil && other.DisableDamage != nil {
		m.DisableDamage.Merge(*other.DisableDamage)
	} else if m.DisableDamage == nil && other.DisableDamage != nil {
		copy := *other.DisableDamage
		m.DisableDamage = &copy
	}
}
