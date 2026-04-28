package lootables

func (l *Lootables) Merge(other Lootables) {
	l.Fills = append(l.Fills, other.Fills...)
	l.Loots = append(l.Loots, other.Loots...)
	l.If = append(l.If, other.If...)
	l.Unless = append(l.Unless, other.Unless...)
}
