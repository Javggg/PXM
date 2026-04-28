package projectiles

func (p *Projectiles) Merge(other Projectiles) {
	p.Projectiles = append(p.Projectiles, other.Projectiles...)
	p.If = append(p.If, other.If...)
	p.Unless = append(p.Unless, other.Unless...)
}
