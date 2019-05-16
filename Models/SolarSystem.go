package Models

import (
	"../Utils"
)

type SolarSystem struct {
 	planets 			map[int] *Planet 
 	stardate 			int `default:0`
 	maxRainIntensity 	float64 `default:-1`
 	dayMaxRainIntensity	int `default:-1`
}

func (ss *SolarSystem) Initialize(){
	ss.planets = make(map[int] *Planet)
}

func (ss *SolarSystem) AddPlanet(p *Planet){
	ss.planets[len(ss.planets)] = p
}

func(ss *SolarSystem) MoveOneDay(){
	ss.stardate++
	for i:=0; i<len(ss.planets); i++ {
		ss.planets[i].Move()
	}
}

func(ss *SolarSystem) GetDayMaxRainIntensity() int{
	return ss.dayMaxRainIntensity
}

func(ss *SolarSystem) GetForecast() Forecast{
	var fc Forecast

	if len(ss.planets)!=3{
		panic("El sistema solar debe tener 3 planetas")
	}

	fc.Drought = ss.drought()
	fc.OptimalWeather = ss.optimalWeather()
	fc.Rainy = ss.rainy()
	
	if fc.Rainy{
		intensity := ss.getRainIntensity()

		if ss.maxRainIntensity < intensity{
			ss.maxRainIntensity = intensity
			ss.dayMaxRainIntensity = ss.stardate
		}
	}

	return fc
}

func (ss *SolarSystem) drought() bool{
	p0 := ss.planets[0].GetCartesianPosition()
	p1 := ss.planets[1].GetCartesianPosition()
	p2 := ss.planets[2].GetCartesianPosition()
	sun := CartesianPosition{0,0}


	triangle := Triangle{p0, p1, p2}
	if triangle.IsTriangle() {
		return false
		
	}

	triangle.CVertex = &sun
	if triangle.IsTriangle() {
		return false
	}

	return true
}

func (ss * SolarSystem) rainy() bool{
	p0 := ss.planets[0].GetCartesianPosition()
	p1 := ss.planets[1].GetCartesianPosition()
	p2 := ss.planets[2].GetCartesianPosition()
	sun := CartesianPosition{0,0}

	triangleX := Triangle{p0, p1, p2}
	triangle0 := Triangle{&sun, p1, p2}
	triangle1 := Triangle{p0, &sun, p2}
	triangle2 := Triangle{p0, p1, &sun}

	return triangleX.Area()>0.0001 && Utils.EqualsFloat(triangleX.Area(), triangle0.Area() + triangle1.Area() + triangle2.Area())
}

func (ss *SolarSystem) getRainIntensity() float64{
	if ss.rainy() {
		p0 := ss.planets[0].GetCartesianPosition()
		p1 := ss.planets[1].GetCartesianPosition()
		p2 := ss.planets[2].GetCartesianPosition()

		triangleX := Triangle{p0, p1, p2}

		return triangleX.Perimeter()
	}

	return -1
}

func (ss *SolarSystem) optimalWeather() bool{
	p0 := ss.planets[0].GetCartesianPosition()
	p1 := ss.planets[1].GetCartesianPosition()
	p2 := ss.planets[2].GetCartesianPosition()
	sun := CartesianPosition{0,0}

	triangle := Triangle{p0, p1, p2}
	if triangle.IsTriangle() {
		return false
		
	}

	triangle.CVertex = &sun
	return triangle.IsTriangle()
}

