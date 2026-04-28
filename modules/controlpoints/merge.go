package controlpoints

func (c *ControlPoints) Merge(other ControlPoints) {
	c.ControlPoints = append(c.ControlPoints, other.ControlPoints...)
	c.If = append(c.If, other.If...)
	c.Unless = append(c.Unless, other.Unless...)
}
