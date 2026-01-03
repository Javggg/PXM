package info

func (a *Authors) Merge(other Authors) {
	a.Authors = append(a.Authors, other.Authors...)
}

func (c *Contributors) Merge(other Contributors) {
	c.Contributors = append(c.Contributors, other.Contributors...)
}
