package cores

func (c *Cores) Merge(other Cores) {
	c.Cores = append(c.Cores, other.Cores...)
}
