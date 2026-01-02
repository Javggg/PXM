package killrewards

func (kr *KillRewards) Merge(other KillRewards) {
	kr.KillRewards = append(kr.KillRewards, other.KillRewards...)
}
