package worldborders

func (w *WorldBorders) Merge(other WorldBorders) {
	w.WorldBorders = append(w.WorldBorders, other.WorldBorders...)
	w.If = append(w.If, other.If...)
	w.Unless = append(w.Unless, other.Unless...)
}
