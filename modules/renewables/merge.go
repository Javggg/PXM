package renewables

func (r *Renewables) Merge(other Renewables) {
	r.Renewables = append(r.Renewables, other.Renewables...)
	r.If = append(r.If, other.If...)
	r.Unless = append(r.Unless, other.Unless...)
}
