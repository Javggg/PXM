package authors

func (a *Authors) Merge(other Authors) {
	a.Authors = append(a.Authors, other.Authors...)
}
