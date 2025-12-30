package maproot

func (m *Map) Merge(other Map) {
	m.Authors.Merge(*other.Authors)
	m.Regions.Merge(*other.Regions)
}
