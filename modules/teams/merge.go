package teams

func (t *Teams) Merge(other Teams) {
	t.Teams = append(t.Teams, other.Teams...)
}
