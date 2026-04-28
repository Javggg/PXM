package regions

func (r *Regions) Merge(other Regions) {
	r.RegionContainer.Merge(other.RegionContainer)
	r.Applies = append(r.Applies, other.Applies...)
}

func (rc *RegionContainer) Merge(other RegionContainer) {
	rc.RegionRefs = append(rc.RegionRefs, other.RegionRefs...)

	rc.Unions = append(rc.Unions, other.Unions...)
	rc.Negatives = append(rc.Negatives, other.Negatives...)
	rc.Complements = append(rc.Complements, other.Complements...)
	rc.Intersects = append(rc.Intersects, other.Intersects...)
	rc.Translates = append(rc.Translates, other.Translates...)
	rc.Mirrors = append(rc.Mirrors, other.Mirrors...)

	rc.Cuboids = append(rc.Cuboids, other.Cuboids...)
	rc.Cylinders = append(rc.Cylinders, other.Cylinders...)
	rc.Spheres = append(rc.Spheres, other.Spheres...)
	rc.Points = append(rc.Points, other.Points...)
	rc.Rectangles = append(rc.Rectangles, other.Rectangles...)
	rc.Circles = append(rc.Circles, other.Circles...)
	rc.Blocks = append(rc.Blocks, other.Blocks...)

	rc.Halves = append(rc.Halves, other.Halves...)
	rc.Below = append(rc.Below, other.Below...)
	rc.Above = append(rc.Above, other.Above...)

	rc.Empty = append(rc.Empty, other.Empty...)
	rc.Nowhere = append(rc.Nowhere, other.Nowhere...)
	rc.Everywhere = append(rc.Everywhere, other.Everywhere...)
	rc.If = append(rc.If, other.If...)
	rc.Unless = append(rc.Unless, other.Unless...)
}
