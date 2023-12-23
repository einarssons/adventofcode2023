package utils

type Pos2D struct {
	Row, Col int
}

func (p Pos2D) Neg() Pos2D {
	return Pos2D{-p.Row, -p.Col}
}

func (p Pos2D) Sub(other Pos2D) Pos2D {
	return Pos2D{p.Row - other.Row, p.Col - other.Col}
}

func (p Pos2D) Add(other Pos2D) Pos2D {
	return Pos2D{p.Row + other.Row, p.Col + other.Col}
}

func (p Pos2D) Mul(factor int) Pos2D {
	return Pos2D{factor * p.Row, factor * p.Col}
}

// left is a 90 degree left rotation of the vector
func (p Pos2D) Left() Pos2D {
	return Pos2D{-p.Col, p.Row}
}

// right is a 90 degree right rotation of the vector
func (p Pos2D) Right() Pos2D {
	return Pos2D{p.Col, -p.Row}
}
