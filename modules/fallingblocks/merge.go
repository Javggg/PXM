package fallingblocks

func (fb *FallingBlocks) Merge(other FallingBlocks) {
	fb.Rules = append(fb.Rules, other.Rules...)
	fb.If = append(fb.If, other.If...)
	fb.Unless = append(fb.Unless, other.Unless...)
}
