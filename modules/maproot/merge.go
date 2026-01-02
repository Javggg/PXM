package maproot

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

	// Replacements
	if m.Replacements != nil && other.Replacements != nil {
		m.Replacements.Merge(*other.Replacements)
	} else if m.Replacements == nil && other.Replacements != nil {
		copy := *other.Replacements
		m.Replacements = &copy
	}

	// Actions
	if m.Actions != nil && other.Actions != nil {
		m.Actions.Merge(*other.Actions)
	} else if m.Actions == nil && other.Actions != nil {
		copy := *other.Actions
		m.Actions = &copy
	}
}
