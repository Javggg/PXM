package constants

func (c *Constants) Merge(other Constants) {
	c.Constants = append(c.Constants, other.Constants...)
}
