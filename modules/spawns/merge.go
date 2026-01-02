package spawns

func (s *Spawns) Merge(other Spawns) {
	s.Spawns = append(s.Spawns, other.Spawns...)
}

func (r *Respawns) Merge(other Respawns) {
	r.Respawns = append(r.Respawns, other.Respawns...)
}
