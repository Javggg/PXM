package spawns

func (s *Spawns) Merge(other Spawns) {
	s.Spawns = append(s.Spawns, other.Spawns...)
	s.If = append(s.If, other.If...)
	s.Unless = append(s.Unless, other.Unless...)
}

func (r *Respawns) Merge(other Respawns) {
	r.Respawns = append(r.Respawns, other.Respawns...)
	r.If = append(r.If, other.If...)
	r.Unless = append(r.Unless, other.Unless...)
}
