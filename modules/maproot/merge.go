package maproot

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

	// Regions
	if m.Regions != nil && other.Regions != nil {
		m.Regions.Merge(*other.Regions)
	} else if m.Regions == nil && other.Regions != nil {
		copy := *other.Regions
		m.Regions = &copy
	}
}
