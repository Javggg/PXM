package shops

func (s *Shops) Merge(other Shops) {
	s.Shops = append(s.Shops, other.Shops...)
}

func (sk *Shopkeepers) Merge(other Shopkeepers) {
	sk.Shopkeepers = append(sk.Shopkeepers, other.Shopkeepers...)
}
