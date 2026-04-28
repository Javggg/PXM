package constants

func (c *Constants) Merge(other Constants) {
	c.Constants = append(c.Constants, other.Constants...)
	c.If = append(c.If, other.If...)
	c.Unless = append(c.Unless, other.Unless...)
}
