package actions

func (a *Actions) Merge(other Actions) {
	a.ActionContainer.Merge(other.ActionContainer)
}

func (ac *ActionContainer) Merge(other ActionContainer) {
	ac.Action = append(ac.Action, other.Action...)
	ac.Trigger = append(ac.Trigger, other.Trigger...)
}
