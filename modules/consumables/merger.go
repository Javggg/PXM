package consumables

func (c *Consumables) Merge(other Consumables) {
	c.Consumables = append(c.Consumables, other.Consumables...)
}
