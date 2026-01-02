package spawners

func (s *Spawners) Merge(other Spawners) {
	s.Spawners = append(s.Spawners, other.Spawners...)
}
