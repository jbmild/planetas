package Modelos

import (
	"math"
)

type Planeta struct {
  	Distancia   float64 //En KM
  	Velocidad	float64 //En grados
  	AnguloActual float64 //En grados
}

func (p *Planeta) PosicionCartesiana() PosicionCartesiana{
	var fAng float64 = float64(int64(p.AnguloActual)%360)
	var anguloRad float64 = fAng / 180

	var pos PosicionCartesiana
	pos.X = p.Distancia * math.Sin(anguloRad * math.Pi)
	pos.Y = p.Distancia * math.Cos(anguloRad * math.Pi)

	return pos
}

func (p *Planeta) DesplazarDias(d int64){
	p.AnguloActual += p.Velocidad * float64(d)
}

