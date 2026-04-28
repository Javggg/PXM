package itemmods

func (im *ItemMods) Merge(other ItemMods) {
	im.Rules = append(im.Rules, other.Rules...)
	im.If = append(im.If, other.If...)
	im.Unless = append(im.Unless, other.Unless...)
}
