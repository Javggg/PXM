package damage

func (dd *DisableDamage) Merge(other DisableDamage) {
	dd.Damage = append(dd.Damage, other.Damage...)
	dd.If = append(dd.If, other.If...)
	dd.Unless = append(dd.Unless, other.Unless...)
}
