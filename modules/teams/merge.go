package teams

func (t *Teams) Merge(other Teams) {
	t.Teams = append(t.Teams, other.Teams...)
	t.If = append(t.If, other.If...)
	t.Unless = append(t.Unless, other.Unless...)
}
