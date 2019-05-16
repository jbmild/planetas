package Modelos

import (
	"math"
)

type SistemaSolar struct {
 	planetas map[string] *Planeta 
 	fechaEstelar int64 `default:0`
 	recta *Recta
}

func (ss *SistemaSolar) Inicializar(){
	ss.recta = nil
	ss.planetas = make(map[string] *Planeta)

	ss.planetas["Vulcano"] = &Planeta{1000, 5, 0}
	ss.planetas["Ferengi"] = &Planeta{500, -1, 0}
	ss.planetas["Betasoide"] = &Planeta{2000, -3, 0}

	ss.SetFechaEstelar(0)
}

func (ss *SistemaSolar) GetFechaEstelar() int64{
	return ss.fechaEstelar
}

func (ss *SistemaSolar) SetFechaEstelar(fecha int64){
	diferencia := fecha - ss.fechaEstelar
	ss.fechaEstelar = fecha
	ss.recta = nil

	ss.planetas["Vulcano"].DesplazarDias(diferencia)
	ss.planetas["Ferengi"].DesplazarDias(diferencia)
	ss.planetas["Betasoide"].DesplazarDias(diferencia)
}

func (ss *SistemaSolar) HaySequia() bool{
	if ss.recta == nil{
		var recta Recta = ss.planetas["Vulcano"].PosicionCartesiana().obtenerRecta(ss.planetas["Ferengi"].PosicionCartesiana())
		ss.recta = &recta
	}

	return ss.recta.Pertenece(ss.planetas["Betasoide"].PosicionCartesiana()) && ss.recta.Pertenece(PosicionCartesiana{0,0})
}

func(ss *SistemaSolar) HayLluvia() bool{
	var p0 PosicionCartesiana = ss.planetas["Vulcano"].PosicionCartesiana()
	var p1 PosicionCartesiana = ss.planetas["Ferengi"].PosicionCartesiana()
	var p2 PosicionCartesiana = ss.planetas["Betasoide"].PosicionCartesiana()

	px := PosicionCartesiana{0,0}

	trinagulox := Triangulo{p0, p1, p2}
	triangulo0 := Triangulo{p1, p2, px}
	triangulo1 := Triangulo{p0, p2, px}
	triangulo2 := Triangulo{p0, p1, px}

	return RedondearPrecision4(trinagulox.Area()) == RedondearPrecision4(triangulo0.Area() + triangulo1.Area() + triangulo2.Area()) && trinagulox.Area()>0
}

func(ss *SistemaSolar) HayCondicionesOptimas() bool{
	if ss.recta == nil{
		var recta Recta = ss.planetas["Vulcano"].PosicionCartesiana().obtenerRecta(ss.planetas["Ferengi"].PosicionCartesiana())
		ss.recta = &recta
	}
	return ss.recta.Pertenece(ss.planetas["Betasoide"].PosicionCartesiana()) && !ss.recta.Pertenece(PosicionCartesiana{0,0})
}


func RedondearPrecision4(x float64) float64{
	return math.Floor(x*10000)/10000
}

func Abs(x float64) float64{
	if x < 0{
		return -x
	}

	return x
}