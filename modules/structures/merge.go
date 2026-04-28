package structures

func (s *Structures) Merge(other Structures) {
	s.Structures = append(s.Structures, other.Structures...)
	s.Dynamics = append(s.Dynamics, other.Dynamics...)
	s.If = append(s.If, other.If...)
	s.Unless = append(s.Unless, other.Unless...)
}
