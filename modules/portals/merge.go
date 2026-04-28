package portals

func (p *Portals) Merge(other Portals) {
	p.Portals = append(p.Portals, other.Portals...)
	p.If = append(p.If, other.If...)
	p.Unless = append(p.Unless, other.Unless...)
}
