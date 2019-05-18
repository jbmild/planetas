package main

import (
	"./Models"
	"github.com/gin-gonic/gin"
   	_ "github.com/go-sql-driver/mysql"
   	"strconv"
   	"./Utils"
   	"fmt"
   	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Bienvenido al pronóstico de la universo. ¿Como podemos ayudarte?",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/clima", answerWeatherOfDay)
	r.GET("/procesar10anios", processJob)
	r.GET("/vaciarbasededatos", truncateTable)
	r.Run(":8080")
}

func answerWeatherOfDay(c *gin.Context) {
	q := c.Request.URL.Query()
  	dia := q["dia"][0]

	day, err := strconv.Atoi(dia)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "El día ingresado es invalido"})  
		return
    }
    if day<0{
    	c.JSON(http.StatusBadRequest, gin.H{"message": "El día ingresado es invalido"})  
		return
    }

    //day := 5

	db := Utils.GetConnection("root", "glamit10", "127.0.0.1", "3306", "planets")

	query := fmt.Sprintf("SELECT day, drought, optimal_weather, rainy FROM forecasts WHERE day=%d", day)
	rows, err := db.Query(query)
    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Ocurrió un error al procesar su petición."})       
		return
    }

	for rows.Next() {
        var forecast Models.Forecast
        if err := rows.Scan(&forecast.Day, &forecast.Drought, &forecast.OptimalWeather, &forecast.Rainy); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Ocurrió un error al procesar su petición."})       
			return
        }

		c.JSON(http.StatusOK, gin.H{"dia": forecast.Day, "clima": forecast.ToString()})        
		return
    }

	c.JSON(http.StatusNotFound, gin.H{"message": "Lo sentimos, no tenemos disponible el día solicitado."})
}

func processJob(c *gin.Context) {
	go calculate10Years()
	c.JSON(http.StatusOK, gin.H{"message": "La tarea se inicio correctamente."})

}

func calculate10Years(){
	years := 10 //Amount of years to calculate
	dayFractionCount := 1 //This is set to seconds, but it could be any fraction qty

	//Connect to database
	db := Utils.GetConnection("root", "glamit10", "127.0.0.1", "3306", "planets")

	var ss Models.SolarSystem
	ss.Initialize(dayFractionCount, false)

	ss.AddPlanet(&Models.Planet{"Vulcano", 1000, 5, 0})
	ss.AddPlanet(&Models.Planet{"Ferengi", 500, -1, 0})
	ss.AddPlanet(&Models.Planet{"Betasoide", 2000, -3, 0})

	for d:=1; d<=(365*years); d++{
	 	fc := ss.GetForecastOfNextDay()

	 	t := fmt.Sprintf("INSERT INTO forecasts (day, drought, optimal_weather, rainy) VALUES (%d, %d, %d, %d)", fc.Day, fc.Drought, fc.OptimalWeather, fc.Rainy)
        if !Utils.ExecuteQuery(db, t){
        	return
        }
	}	


    db.Close()
}

func truncateTable(c *gin.Context){
	db := Utils.GetConnection("root", "glamit10", "127.0.0.1", "3306", "planets")

    //Create table
    Utils.ExecuteQuery(db, "TRUNCATE TABLE forecasts")
    c.JSON(http.StatusOK, gin.H{"message": "La información se borro correctamente."})
}
