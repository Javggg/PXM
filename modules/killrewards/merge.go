package killrewards

func (kr *KillRewards) Merge(other KillRewards) {
	kr.KillRewards = append(kr.KillRewards, other.KillRewards...)
	kr.If = append(kr.If, other.If...)
	kr.Unless = append(kr.Unless, other.Unless...)
}
