package actions

func (a *Actions) Merge(other Actions) {
	a.Actions = append(a.Actions, other.Actions...)
	a.Triggers = append(a.Triggers, other.Triggers...)
	a.If = append(a.If, other.If...)
	a.Unless = append(a.Unless, other.Unless...)
}
