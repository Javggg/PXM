package spawners

func (s *Spawners) Merge(other Spawners) {
	s.Spawners = append(s.Spawners, other.Spawners...)
	s.If = append(s.If, other.If...)
	s.Unless = append(s.Unless, other.Unless...)
}
