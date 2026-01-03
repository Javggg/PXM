package modes

func (m *Modes) Merge(other Modes) {
	m.Modes = append(m.Modes, other.Modes...)
}
