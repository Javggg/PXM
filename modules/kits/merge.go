package kits

func (k *Kits) Merge(other Kits) {
	k.Kits = append(k.Kits, other.Kits...)
	k.Give = append(k.Give, other.Give...)
	k.Take = append(k.Take, other.Take...)
	k.Lend = append(k.Lend, other.Lend...)
	k.If = append(k.If, other.If...)
	k.Unless = append(k.Unless, other.Unless...)
}
