package wools

func (w *Wools) Merge(other Wools) {
	w.Wools = append(w.Wools, other.Wools...)
}
