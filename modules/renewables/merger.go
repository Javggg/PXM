package renewables

func (r *Renewables) Merge(other Renewables) {
	r.Renewables = append(r.Renewables, other.Renewables...)
}
