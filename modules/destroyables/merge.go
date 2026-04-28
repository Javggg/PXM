package destroyables

func (d *Destroyables) Merge(other Destroyables) {
	d.Destroyables = append(d.Destroyables, other.Destroyables...)
	d.If = append(d.If, other.If...)
	d.Unless = append(d.Unless, other.Unless...)
}
