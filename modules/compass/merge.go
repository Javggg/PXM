package compass

func (c *Compass) Merge(other Compass) {
	c.Players = append(c.Players, other.Players...)
	c.Flags = append(c.Flags, other.Flags...)
	c.If = append(c.If, other.If...)
	c.Unless = append(c.Unless, other.Unless...)
}
