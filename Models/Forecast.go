package Models

type Forecast struct {
  	Drought			bool `default:false`
  	OptimalWeather	bool `default:false`
  	Rainy 			bool `default:false`
}