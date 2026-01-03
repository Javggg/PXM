package damage

func (dd *DisableDamage) Merge(other DisableDamage) {
	dd.Damage = append(dd.Damage, other.Damage...)
}
