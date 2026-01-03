package worldborders

func (w *WorldBorders) Merge(other WorldBorders) {
	w.WorldBorders = append(w.WorldBorders, other.WorldBorders...)
}
