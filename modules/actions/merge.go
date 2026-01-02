package actions

func (a *Actions) Merge(other Actions) {
	a.Actions = append(a.Actions, other.Actions...)
	a.Triggers = append(a.Triggers, other.Triggers...)
}
