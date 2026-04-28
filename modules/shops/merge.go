package shops

func (s *Shops) Merge(other Shops) {
	s.Shops = append(s.Shops, other.Shops...)
	s.If = append(s.If, other.If...)
	s.Unless = append(s.Unless, other.Unless...)
}

func (sk *Shopkeepers) Merge(other Shopkeepers) {
	sk.Shopkeepers = append(sk.Shopkeepers, other.Shopkeepers...)
	sk.If = append(sk.If, other.If...)
	sk.Unless = append(sk.Unless, other.Unless...)
}
