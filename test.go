package main

import (
	"fmt"
	"net/http"
	"strconv"

	"project/test/mathQuestion"

	"github.com/gin-gonic/gin"
)

type sales struct {
	Year     int    `json:"year"`
	Vehicle  string `json:"vehicle"`
	Quantity int    `json:"quantity"`
}

var vehicleSales = []sales{
	{Year: 2021, Vehicle: "Elise", Quantity: 200},
	{Year: 2022, Vehicle: "Elise", Quantity: 130},
	{Year: 2023, Vehicle: "Elise", Quantity: 80},
}

func main() {
	http := gin.Default()
	http.GET("/vehicle-sales", getSales)
	http.GET("/vehicle-sales/:year", getSalesByYear)
	http.POST("/vehicle-sales", postSales) // Don't know how to function
	http.Run("localhost:8080")
}

// respond with a list of all sales as JSON
func getSales(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, vehicleSales)
}

// respond with a list of sales for a year as JSON
func getSalesByYear(c *gin.Context) {
	year := c.Param("year")
	var vehicleSales_year []sales

	for _, s := range vehicleSales {
		if year == strconv.Itoa(s.Year) {
			vehicleSales_year = append(vehicleSales_year, s)
		}
	}

	c.IndentedJSON(http.StatusOK, vehicleSales_year)
}

// add a sales from JSON received in request body
func postSales(c *gin.Context) {
	var newSales sales // = sales{Year: 2019, Vehicle: "Elise", Quantity: 110}

	if err := c.BindJSON(&newSales); err != nil {
		return
	}

	vehicleSales = append(vehicleSales, newSales)
	c.IndentedJSON(http.StatusCreated, newSales)

}

func QandA() {
	q := mathQuestion.RandomQuestion()
	fmt.Println("Enter -1 to exit.")
	fmt.Printf("%v %c %v = ", q.Num1, q.Operator, q.Num2)

	var ansString string
	fmt.Scanln(&ansString)

	ans, err := strconv.Atoi(ansString)
	if err == nil {
		if ans == -1 {
			return
		}
		fmt.Printf("%s\n", mathQuestion.CheckAns(ans, q))
	} else {
		fmt.Printf("\033[31mNot a number\033[0m\n\n")
	}
	QandA()
}
