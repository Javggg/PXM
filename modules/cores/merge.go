package cores

func (c *Cores) Merge(other Cores) {
	c.Cores = append(c.Cores, other.Cores...)
	c.If = append(c.If, other.If...)
	c.Unless = append(c.Unless, other.Unless...)
}
