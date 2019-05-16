package Modelos

type Recta struct {
  Pendiente float64
  OrdenadaAlOrigen float64
}

func (r *Recta) Pertenece(p PosicionCartesiana) bool{
	if r.Pendiente == 0 && RedondearPrecision4(p.X)==0{
		return true
	}

	if RedondearPrecision4(p.Y) == RedondearPrecision4(r.calcular(p.X)){
		return true
	} else {
		return false
	}
}

func (r *Recta) calcular(x float64) float64{
	if r.Pendiente != 0{
		return r.Pendiente * x + r.OrdenadaAlOrigen
	} else {
		return r.OrdenadaAlOrigen
	}
}