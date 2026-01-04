package structures

func (s *Structures) Merge(other Structures) {
	s.Structures = append(s.Structures, other.Structures...)
	s.Dynamics = append(s.Dynamics, other.Dynamics...)
}
