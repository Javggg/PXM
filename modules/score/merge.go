package score

func (s *Score) Merge(other Score) {
	s.Boxes = append(s.Boxes, other.Boxes...)
	s.If = append(s.If, other.If...)
	s.Unless = append(s.Unless, other.Unless...)
}
