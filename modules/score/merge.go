package score

func (s *Score) Merge(other Score) {
	s.Boxes = append(s.Boxes, other.Boxes...)
}
