package projectiles

func (p *Projectiles) Merge(other Projectiles) {
	p.Projectiles = append(p.Projectiles, other.Projectiles...)
}
