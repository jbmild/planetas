package main

import (
	"./Models"
    "database/sql"
   	_ "github.com/go-sql-driver/mysql"
   	"fmt"
)

func main() {
	years := 10 //Amount of years to calculate
	dayFractionCount := 1 //This is set to seconds, but it could be any fraction qty

	//Connect to database
	db, err := sql.Open("mysql", "root:glamit10@tcp(127.0.0.1:3306)/planets")
	if err != nil {
        panic(err.Error())
    }

    defer db.Close()

    //Drop table
    drop, err := db.Query("DROP TABLE IF EXISTS forecasts")
    if err != nil {
        panic(err.Error())
    }
    drop.Close()

    //Create table
    create, err := db.Query("CREATE TABLE forecasts (" +
                                "day int(10) unsigned NOT NULL," +
                                "drought int(10) unsigned NOT NULL," +
                                "optimal_weather int(10) unsigned NOT NULL," +
                                "rainy int(10) unsigned NOT NULL," +
                                "UNIQUE KEY day_unique (day)" +
                              ") ENGINE=InnoDB")
    if err != nil {
        panic(err.Error())
    }
    create.Close()

	var ss Models.SolarSystem
	ss.Initialize(dayFractionCount, false)

	ss.AddPlanet(&Models.Planet{"Vulcano", 1000, 5, 0})
	ss.AddPlanet(&Models.Planet{"Ferengi", 500, -1, 0})
	ss.AddPlanet(&Models.Planet{"Betasoide", 2000, -3, 0})

	for d:=1; d<=(365*years); d++{
	 	fc := ss.GetForecastOfNextDay()

	 	t := fmt.Sprintf("INSERT INTO forecasts (day, drought, optimal_weather, rainy) VALUES (%d, %d, %d, %d)", fc.Day, fc.Drought, fc.OptimalWeather, fc.Rainy)

	 	insert, err := db.Query(t)
	    if err != nil {
	        panic(err.Error())
	    }
	    insert.Close()
	}	
}