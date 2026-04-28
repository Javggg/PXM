package modes

func (m *Modes) Merge(other Modes) {
	m.Modes = append(m.Modes, other.Modes...)
	m.If = append(m.If, other.If...)
	m.Unless = append(m.Unless, other.Unless...)
}
