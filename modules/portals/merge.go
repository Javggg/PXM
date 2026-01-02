package portals

func (p *Portals) Merge(other Portals) {
	p.Portals = append(p.Portals, other.Portals...)
}
