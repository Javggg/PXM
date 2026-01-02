package destroyables

func (d *Destroyables) Merge(other Destroyables) {
	d.Destroyables = append(d.Destroyables, other.Destroyables...)
}
