package Models

import (
	"math"
	"../Utils"
)

type Planet struct {
	Name					string
  	DistanceFromTheSun   	float64 //En KM
  	AngularSpeed			float64 //En grados
  	AnglePosition 			float64 //En grados
}

func (p *Planet) GetCartesianPosition() *CartesianPosition{
	var radAngle float64 = p.AnglePosition * math.Pi / 180

	var pos CartesianPosition
	pos.X = p.DistanceFromTheSun * math.Cos(radAngle)
	pos.Y = p.DistanceFromTheSun * math.Sin(radAngle)

	return &pos
}

func (p *Planet) MoveOneFraction(dayFractionCount int){

	p.AnglePosition += p.AngularSpeed / float64(dayFractionCount)
	if Utils.Abs(p.AnglePosition) >= 360{
		if p.AnglePosition > 0{
			p.AnglePosition -= 360
		}else{
			p.AnglePosition += 360
		}
	}
}

