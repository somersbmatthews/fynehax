package r2

// Box defines a box in R2
//
//                A
//                |
//                |
//                |
//                v
//    (1) A.X,A.Y +------+ A.X+S.X,A.Y (2)
//                |\     |
//                | \    |
//                |  \   |
//                |   \S |
//                |    \ |
//                |     \|
// (3)A.X,A.Y+S.Y +------+ A.X+S.X,A.Y+S.Y (4)
//
type Box struct {

	// A defines the top-left corner of the box
	A Vec2

	// S defines the size of the box
	S Vec2
}

func MakeBox(a, s Vec2) Box {
	return Box{
		A: a,
		S: s,
	}
}

func (b Box) Area() float64 {
	return b.S.X * b.S.Y
}

// GetCorner1 returns the top left corner of the box
func (b Box) GetCorner1() Vec2 {
	return b.A
}

// GetCorner2 returns the top right corner of the box
func (b Box) GetCorner2() Vec2 {
	return b.A.Add(V2(b.S.X, 0))
}

// GetCorner3 returns the bottom left corner of the box.
func (b Box) GetCorner3() Vec2 {
	return b.A.Add(V2(0, b.S.Y))
}

// GetCorner4 returns the bottom right corner of the box.
func (b Box) GetCorner4() Vec2 {
	return b.A.Add(V2(b.S.X, b.S.Y))
}

// Returns the intersection of the box and the line, and a Boolean indicating
// if the box and vector intersect. If they do not collide, the zero vector is
// returned.
func (b Box) Intersect(l Line) (Vec2, bool) {
	// This is transliterated in part from:
	//
	// https://github.com/JulNadeauCA/libagar/blob/master/gui/primitive.c

	faces := []Line{
		b.Top(),
		b.Left(),
		b.Right(),
		b.Bottom(),
	}

	dists := []float64{-1, -1, -1, -1}
	intersects := []bool{false, false, false, false}
	intersectPoints := make([]Vec2, 4)

	shortest_dist := float64(-1.0)
	best := -1

	for i := range faces {
		in, ok := IntersectLines(faces[i], l)
		if !ok {
			continue
		}
		dists[i] = in.Length()
		intersects[i] = ok
		intersectPoints[i] = in

		if (dists[i] < shortest_dist) || (shortest_dist == float64(-1)) {
			shortest_dist = dists[i]
			best = i
		}
	}

	if shortest_dist < 0 {
		return V2(0, 0), false
	}

	return intersectPoints[best], true
}

// Top returns the top face of the box.
func (b Box) Top() Line {
	return MakeLineFromEndpoints(b.GetCorner1(), b.GetCorner2())
}

// Left returns the left face of the box.
func (b Box) Left() Line {
	return MakeLineFromEndpoints(b.GetCorner1(), b.GetCorner3())
}

// Right returns the right face of the box.
func (b Box) Right() Line {
	return MakeLineFromEndpoints(b.GetCorner2(), b.GetCorner4())
}

// Bottom returns the bottom face of the box.
func (b Box) Bottom() Line {
	return MakeLineFromEndpoints(b.GetCorner3(), b.GetCorner4())
}

func (b Box) Center() Vec2 {
	return b.A.Add(b.S.Scale(0.5))
}
