package Modelos

type Triangulo struct {
 	P0 PosicionCartesiana
 	P1 PosicionCartesiana
 	P2 PosicionCartesiana
}

func(t *Triangulo) Area() float64{
	return 0.5 *(-t.P1.Y * t.P2.X + t.P0.Y * (-t.P1.X + t.P2.X) + t.P0.X * (t.P1.Y - t.P2.Y) + t.P1.X * t.P2.Y);
}