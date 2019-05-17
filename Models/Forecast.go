package Models

type Forecast struct {
  	Drought			int `default:0`
  	OptimalWeather	int `default:0`
  	Rainy 			int `default:0`
  	Day				int `default:0`
}

func(f1 *Forecast) Equals(f2 Forecast) bool{
	return (f1.Drought-f2.Drought)==0 && (f1.OptimalWeather-f2.OptimalWeather)==0 && (f1.Rainy-f2.Rainy)==0
}