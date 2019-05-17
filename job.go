package main

import (
	"./Models"
   	"fmt"
    "./Utils"
)

func main() {
	years := 10 //Amount of years to calculate
	dayFractionCount := 1 //This is set to seconds, but it could be any fraction qty

	//Connect to database
	db := Utils.GetConnection("root", "glamit10", "127.0.0.1", "3306", "planets")



    //Drop table
    Utils.ExecuteQuery(db, "DROP TABLE IF EXISTS forecasts")

    //Create table
    Utils.ExecuteQuery(db, "CREATE TABLE forecasts (" +
                                "day int(10) unsigned NOT NULL," +
                                "drought int(10) unsigned NOT NULL," +
                                "optimal_weather int(10) unsigned NOT NULL," +
                                "rainy int(10) unsigned NOT NULL," +
                                "UNIQUE KEY day_unique (day)" +
                              ") ENGINE=InnoDB")

	var ss Models.SolarSystem
	ss.Initialize(dayFractionCount, false)

	ss.AddPlanet(&Models.Planet{"Vulcano", 1000, 5, 0})
	ss.AddPlanet(&Models.Planet{"Ferengi", 500, -1, 0})
	ss.AddPlanet(&Models.Planet{"Betasoide", 2000, -3, 0})

	for d:=1; d<=(365*years); d++{
	 	fc := ss.GetForecastOfNextDay()

	 	t := fmt.Sprintf("INSERT INTO forecasts (day, drought, optimal_weather, rainy) VALUES (%d, %d, %d, %d)", fc.Day, fc.Drought, fc.OptimalWeather, fc.Rainy)
        Utils.ExecuteQuery(db, t)
	}	


    db.Close()
}