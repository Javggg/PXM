package fallingblocks

func (fb *FallingBlocks) Merge(other FallingBlocks) {
	fb.Rules = append(fb.Rules, other.Rules...)
}
