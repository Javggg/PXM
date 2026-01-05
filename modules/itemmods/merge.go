package itemmods

func (im *ItemMods) Merge(other ItemMods) {
	im.Rules = append(im.Rules, other.Rules...)
}
