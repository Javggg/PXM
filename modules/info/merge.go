package info

func (a *Authors) Merge(other Authors) {
	a.Authors = append(a.Authors, other.Authors...)
	a.If = append(a.If, other.If...)
	a.Unless = append(a.Unless, other.Unless...)
}

func (c *Contributors) Merge(other Contributors) {
	c.Contributors = append(c.Contributors, other.Contributors...)
	c.If = append(c.If, other.If...)
	c.Unless = append(c.Unless, other.Unless...)
}
