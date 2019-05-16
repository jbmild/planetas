package main

import (
	"./Modelos"
	"fmt"
)

func main() {
	anios := 10
	optimo, sequia, lluvia := 0, 0, 0

	var ss Modelos.SistemaSolar
	ss.Inicializar()

	contabilizar(&ss, &optimo, &sequia, &lluvia)
	for d:=1; d<=(365*anios); d++{
	 	ss.SetFechaEstelar(int64(d))
	 	contabilizar(&ss, &optimo, &sequia, &lluvia)
	}

	fmt.Println("Condiciones optimas", optimo)
	fmt.Println("Sequia", sequia)
	fmt.Println("Lluvia", lluvia)
	
}

func contabilizar(ss *Modelos.SistemaSolar, optimo *int, sequia *int, lluvia *int){
	if ss.HayCondicionesOptimas(){
		*optimo++
	}

	if ss.HaySequia(){
		*sequia++
	}

	if ss.HayLluvia(){
		*lluvia++
	}
}