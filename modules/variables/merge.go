package variables

func (v *Variables) Merge(other Variables) {
	v.VariableContainer.Merge(other.VariableContainer)
}

func (vc *VariableContainer) Merge(other VariableContainer) {
	vc.Variable = append(vc.Variable, other.Variable...)
	vc.Array = append(vc.Array, other.Array...)
	vc.WithTeam = append(vc.WithTeam, other.WithTeam...)
	vc.PlayerLocation = append(vc.PlayerLocation, other.PlayerLocation...)
	vc.Score = append(vc.Score, other.Score...)
	vc.TimeLimit = append(vc.TimeLimit, other.TimeLimit...)
	vc.MaxBuildHeight = append(vc.MaxBuildHeight, other.MaxBuildHeight...)
	vc.WorldTime = append(vc.WorldTime, other.WorldTime...)
}
