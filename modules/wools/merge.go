package wools

func (w *Wools) Merge(other Wools) {
	w.Wools = append(w.Wools, other.Wools...)
	w.If = append(w.If, other.If...)
	w.Unless = append(w.Unless, other.Unless...)
}
