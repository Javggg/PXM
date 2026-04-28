package replacements

func (r *Replacements) Merge(other Replacements) {
	r.ReplacementsContainer.Merge(other.ReplacementsContainer)
}

func (rc *ReplacementsContainer) Merge(other ReplacementsContainer) {
	rc.Decimal = append(rc.Decimal, other.Decimal...)
	rc.Player = append(rc.Player, other.Player...)
	rc.Switch = append(rc.Switch, other.Switch...)
	rc.Replacement = append(rc.Replacement, other.Replacement...)
	rc.If = append(rc.If, other.If...)
	rc.Unless = append(rc.Unless, other.Unless...)
}
