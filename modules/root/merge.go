package root

func (m *Map) Merge(other Map) {
	// Constants
	if m.Constants != nil && other.Constants != nil {
		m.Constants.Merge(*other.Constants)
	} else if m.Constants == nil && other.Constants != nil {
		copy := *other.Constants
		m.Constants = &copy
	}

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

	// KillRewards
	if m.KillRewards != nil && other.KillRewards != nil {
		m.KillRewards.Merge(*other.KillRewards)
	} else if m.KillRewards == nil && other.KillRewards != nil {
		copy := *other.KillRewards
		m.KillRewards = &copy
	}
}
