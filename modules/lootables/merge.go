package lootables

func (l *Lootables) Merge(other Lootables) {
	l.Fills = append(l.Fills, other.Fills...)
	l.Loots = append(l.Loots, other.Loots...)
}
