/*package main

import (
	"./Models"
	"fmt"
)

func main() {
	years := 10 //Amount of years to calculate
	dayFractionCount := 1 //This is set to seconds, but it could be any fraction qty

	acDrought, acOptimalWeather, acRainy := 0, 0, 0

	var ss Models.SolarSystem
	ss.Initialize(dayFractionCount, false)

	ss.AddPlanet(&Models.Planet{"Vulcano", 1000, 5, 0})
	ss.AddPlanet(&Models.Planet{"Ferengi", 500, -1, 0})
	ss.AddPlanet(&Models.Planet{"Betasoide", 2000, -3, 0})

	for d:=1; d<=(365*years); d++{
	 	fc := ss.GetForecastOfNextDay()
		acDrought += fc.Drought
		acOptimalWeather += fc.OptimalWeather
		acRainy += fc.Rainy
	}

	fmt.Println("Períodos de sequía", acDrought)
	fmt.Println("Períodos con condiciones óptimas", acOptimalWeather)
	fmt.Println("Períodos de lluvia", acRainy)
	if ss.GetDayMaxRainIntensity() > -1{
		fmt.Println("El dia con mayor lluvia fue el ", ss.GetDayMaxRainIntensity())
	}
	
}*/