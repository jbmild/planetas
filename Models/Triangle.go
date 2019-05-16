package Models

import (
	"../Utils"
)

type Triangle struct {
 	AVertex *CartesianPosition
 	BVertex *CartesianPosition
 	CVertex *CartesianPosition
}

func(t *Triangle) IsTriangle() bool{
	return !Utils.EqualsFloat(t.Area(), 0.0)
}

func(t *Triangle) Area() float64{
	a, b, c := t.AVertex, t.BVertex, t.CVertex

	area := a.X * (b.Y - c.Y) + b.X * (c.Y - a.Y) + c.X * (a.Y - b.Y)
	return Utils.Abs(0.5 * area)
}

func(t *Triangle) Perimeter() float64{
	distanceAB := t.AVertex.Distance(*t.BVertex)
	distanceBC := t.BVertex.Distance(*t.CVertex)
	distanceAC := t.AVertex.Distance(*t.CVertex)

	return distanceAB + distanceBC + distanceAC
}