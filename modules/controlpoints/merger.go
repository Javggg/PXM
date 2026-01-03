package controlpoints

func (c *ControlPoints) Merge(other ControlPoints) {
	c.ControlPoints = append(c.ControlPoints, other.ControlPoints...)
}
