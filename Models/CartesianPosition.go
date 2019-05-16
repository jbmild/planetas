package Models

import (
	"math"
)

type CartesianPosition struct {
	X float64
	Y float64
}

func (position1 *CartesianPosition) Distance(position2 CartesianPosition) float64{
	return math.Sqrt( math.Pow(position1.X - position2.X, 2) + math.Pow(position1.Y - position2.Y, 2) )
}
