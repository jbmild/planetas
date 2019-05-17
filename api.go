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
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/clima", answerWeatherOfDay)
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
	rows := Utils.GetResult(db, query)

	for rows.Next() {
        var forecast Models.Forecast
        if err := rows.Scan(&forecast.Day, &forecast.Drought, &forecast.OptimalWeather, &forecast.Rainy); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Ocurrió un error al procesar su petición.2"})       
			return
        }

		c.JSON(http.StatusOK, gin.H{"dia": forecast.Day, "clima": forecast.ToString()})        
		return
    }
    if err := rows.Err(); err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Ocurrió un error al procesar su petición.1"})       
        return
    }

	c.JSON(http.StatusNotFound, gin.H{"message": "Lo sentimos, no tenemos disponible el día solicitado."})
}
