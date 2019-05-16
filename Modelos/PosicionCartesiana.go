package Modelos

type PosicionCartesiana struct {
  X float64
  Y float64
}

func (pc1 PosicionCartesiana) obtenerRecta(pc2 PosicionCartesiana) Recta{
	var m float64
	var oo float64

	if (pc2.X - pc1.X) != 0 {
		m = (pc2.Y - pc1.Y) / (pc2.X - pc1.X)
	}else{
		m = 0
	}

	oo = pc1.Y - m * pc1.X

	return Recta{m, oo}
}

