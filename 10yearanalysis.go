package main

import (
	"./Models"
	"fmt"
)

func main() {
	years := 10
	acDrought, acOptimalWeather, acRainy := 0, 0, 0

	var ss Models.SolarSystem
	ss.Initialize()

	ss.AddPlanet(&Models.Planet{"Vulcano", 1000, 5, 0})
	ss.AddPlanet(&Models.Planet{"Ferengi", 500, -1, 0})
	ss.AddPlanet(&Models.Planet{"Betasoide", 2000, -3, 0})

	fc := ss.GetForecast()
	if fc.Drought { acDrought++ }
	if fc.OptimalWeather { acOptimalWeather++ }
	if fc.Rainy { acRainy++ }

	for d:=1; d<=(365*years); d++{
		ss.MoveOneDay()
	 	fc := ss.GetForecast()
		if fc.Drought { acDrought++ }
		if fc.OptimalWeather { acOptimalWeather++ }
		if fc.Rainy { acRainy++ }
	}

	fmt.Println("Sequia", acDrought)
	fmt.Println("Condiciones optimas", acOptimalWeather)
	fmt.Println("Lluvia", acRainy)
	if ss.GetDayMaxRainIntensity() > -1{
		fmt.Println("El dia con mayor lluvia fue el ", ss.GetDayMaxRainIntensity())
	}
	
}