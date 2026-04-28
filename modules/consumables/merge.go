package consumables

func (c *Consumables) Merge(other Consumables) {
	c.Consumables = append(c.Consumables, other.Consumables...)
	c.If = append(c.If, other.If...)
	c.Unless = append(c.Unless, other.Unless...)
}
