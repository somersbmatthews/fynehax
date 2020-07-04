package r2

// Line describes a line in R2
//
//  (1) A.X,A.Y  +
//                \
//                 \
//                  \
//                   \
//                    + A.X+S.X,A.Y+S.Y (2)
//
type Line struct {
	// A defines the basis point of the line
	A Vec2

	// S defines the direction and length of the line
	S Vec2
}

func MakeLine(a, s Vec2) Line {
	return Line{
		A: a,
		S: s,
	}
}

// Return a line which has endpoints a, b
func MakeLineFromEndpoints(a, b Vec2) Line {
	s := b.Add(a.Scale(-1))

	return MakeLine(a, s)
}

func (l Line) Endpoint1() Vec2 {
	return l.A
}

func (l Line) Endpoint2() Vec2 {
	return l.A.Add(l.S)
}

// IntersectLines returns the point at which two lines intersect, or the
// zero vector, along with a boolean indicating if the lines intersect.
//
// Based on the code described here: https://stackoverflow.com/a/14795484
func IntersectLines(l1, l2 Line) (Vec2, bool) {
	s10x := l1.Endpoint2().X - l1.Endpoint1().X
	s10y := l1.Endpoint2().Y - l1.Endpoint1().Y
	s32x := l2.Endpoint2().X - l2.Endpoint1().X
	s32y := l2.Endpoint2().Y - l2.Endpoint1().Y

	denom := s10x*s32y - s32x*s10y
	if denom == 0 {
		return V2(0, 0), false
	}
	denomPositive := denom > 0

	s02x := l1.Endpoint1().X - l1.Endpoint1().X
	s02y := l1.Endpoint1().Y - l1.Endpoint1().Y
	s_numer := s10x*s02y - s10y*s02x
	if (s_numer < 0) == denomPositive {
		return V2(0, 0), false
	}

	t_numer := s32x*s02y - s32y*s02x
	if (t_numer < 0) == denomPositive {
		return V2(0, 0), false
	}

	if ((s_numer > denom) == denomPositive) || ((t_numer > denom) == denomPositive) {
		return V2(0, 0), false
	}

	t := t_numer / denom
	return V2(l1.Endpoint1().X+t*s10x, l1.Endpoint1().Y+t*s10y), true
}
