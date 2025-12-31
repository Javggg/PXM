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

	// Regions
	if m.Regions != nil && other.Regions != nil {
		m.Regions.Merge(*other.Regions)
	} else if m.Regions == nil && other.Regions != nil {
		copy := *other.Regions
		m.Regions = &copy
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
}
