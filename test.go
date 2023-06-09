package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/utahta/go-linenotify"

	"net/http"
)

// type Book struct {
// 	ID     string `json:"id"`
// 	Title  string `json:"title"`
// 	Author string `json:"author"`
// }

type ShirtColor struct {
	ID    string `json:"id"`
	Day   string `json:"day"`
	Color string `json:"color"`
}

type Message struct {
	Message string `json:"message"`
}

// JSON-to-Go : https://mholt.github.io/json-to-go/
type WeatherData struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	CurrentWeather       struct {
		Temperature   float64 `json:"temperature"`
		Windspeed     float64 `json:"windspeed"`
		Winddirection float64 `json:"winddirection"`
		Weathercode   int     `json:"weathercode"`
		IsDay         int     `json:"is_day"`
		Time          string  `json:"time"`
	} `json:"current_weather"`
	DailyUnits struct {
		Time                   string `json:"time"`
		Temperature2MMax       string `json:"temperature_2m_max"`
		Temperature2MMin       string `json:"temperature_2m_min"`
		ApparentTemperatureMax string `json:"apparent_temperature_max"`
		ApparentTemperatureMin string `json:"apparent_temperature_min"`
		Sunrise                string `json:"sunrise"`
		Sunset                 string `json:"sunset"`
		UvIndexMax             string `json:"uv_index_max"`
	} `json:"daily_units"`
	Daily struct {
		Time                   []string  `json:"time"`
		Temperature2MMax       []float64 `json:"temperature_2m_max"`
		Temperature2MMin       []float64 `json:"temperature_2m_min"`
		ApparentTemperatureMax []float64 `json:"apparent_temperature_max"`
		ApparentTemperatureMin []float64 `json:"apparent_temperature_min"`
		Sunrise                []string  `json:"sunrise"`
		Sunset                 []string  `json:"sunset"`
		UvIndexMax             []float64 `json:"uv_index_max"`
	} `json:"daily"`
}

// var books = []Book{
// 	{ID: "1", Title: "Harry Potter", Author: "J. K. Rowling"},
// 	{ID: "2", Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
// 	{ID: "3", Title: "The Wizard of Oz", Author: "L. Frank Baum"},
// }

var luckycolor = []ShirtColor{
	{ID: "1", Day: "Sunday", Color: "Red"},
	{ID: "2", Day: "Monday", Color: "Yellow"},
	{ID: "3", Day: "Tuesday", Color: "Pink"},
	{ID: "4", Day: "Wednesday", Color: "Green"},
	{ID: "5", Day: "Thursday", Color: "Orange"},
	{ID: "6", Day: "Friday", Color: "Blue"},
	{ID: "7", Day: "Saturday", Color: "Purple"},
}

var linetoken = "s1P1F7Rm3tO6sxUpAPuY7TuTF315zXav7b0U8QwBoRm"

// func getbookHandler(c *gin.Context) {
// 	c.JSON(
// 		http.StatusOK,
// 		books,
// 	)
// }

// func postbookHandler(c *gin.Context) {
// 	var book Book
// 	if err := c.ShouldBindJSON(&book); err != nil {
// 		c.JSON(
// 			http.StatusBadRequest,
// 			gin.H{"error": err.Error()},
// 		)
// 	}
// 	books = append(books, book)
// 	c.JSON(
// 		http.StatusCreated,
// 		book,
// 	)
// }

// func deletebookHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	for i, a := range books {
// 		if a.ID == id {
// 			books = append(books[:i], books[i+1:]...)
// 		}
// 	}
// 	c.Status(http.StatusNoContent)
// }

func getcolor() string {
	dt := time.Now().Weekday()
	color := "White"
	switch dt.String() {
	case luckycolor[0].Day:
		color = luckycolor[0].Color
		fmt.Println(color)
	case luckycolor[1].Day:
		color = luckycolor[1].Color
		fmt.Println(color)
	case luckycolor[2].Day:
		color = luckycolor[2].Color
		fmt.Println(color)
	case luckycolor[3].Day:
		color = luckycolor[3].Color
		fmt.Println(color)
	case luckycolor[4].Day:
		color = luckycolor[4].Color
		fmt.Println(color)
	case luckycolor[5].Day:
		color = luckycolor[5].Color
		fmt.Println(color)
	case luckycolor[6].Day:
		color = luckycolor[6].Color
		fmt.Println(color)
	}
	return color
}

func getweather() string {
	urlStr := "https://api.open-meteo.com/v1/forecast?latitude=13.75&longitude=100.50&daily=temperature_2m_max,temperature_2m_min,apparent_temperature_max,apparent_temperature_min,sunrise,sunset,uv_index_max&current_weather=true&forecast_days=1&timezone=Asia%2FBangkok"
	res, _ := http.Get(urlStr)
	resBody, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	var data WeatherData
	json.Unmarshal(resBody, &data)
	uvmax := fmt.Sprintf("%v", data.Daily.UvIndexMax[0])
	date := fmt.Sprintf("%v", data.Daily.Time[0])
	tempmax := fmt.Sprintf("%v", data.Daily.Temperature2MMax[0])
	tempmin := fmt.Sprintf("%v", data.Daily.Temperature2MMin[0])
	rise := fmt.Sprintf("%v", data.Daily.Sunrise[0][11:])
	set := fmt.Sprintf("%v", data.Daily.Sunset[0][11:])
	aptempmax := fmt.Sprintf("%v", data.Daily.ApparentTemperatureMax[0])
	aptempmin := fmt.Sprintf("%v", data.Daily.ApparentTemperatureMin[0])
	return "Date : " + date + "\nTemp : " + tempmin + " - " + tempmax + "\nFeel like : " + aptempmin + " - " + aptempmax + "\nSunrise : " + rise + "\nSunset : " + set + "\nUV : " + uvmax
}

func noti(c *gin.Context) {
	dt := time.Now().Weekday()
	// color := "White"
	// switch dt.String() {
	// case luckycolor[0].Day:
	// 	color = luckycolor[0].Color
	// 	fmt.Println(color)
	// case luckycolor[1].Day:
	// 	color = luckycolor[1].Color
	// 	fmt.Println(color)
	// case luckycolor[2].Day:
	// 	color = luckycolor[2].Color
	// 	fmt.Println(color)
	// case luckycolor[3].Day:
	// 	color = luckycolor[3].Color
	// 	fmt.Println(color)
	// case luckycolor[4].Day:
	// 	color = luckycolor[4].Color
	// 	fmt.Println(color)
	// case luckycolor[5].Day:
	// 	color = luckycolor[5].Color
	// 	fmt.Println(color)
	// case luckycolor[6].Day:
	// 	color = luckycolor[6].Color
	// 	fmt.Println(color)
	// }

	color := getcolor()

	weather := getweather()
	print(weather)

	cli := linenotify.NewClient()
	cli.Notify(context.Background(), linetoken, dt.String()+"\nLucky Color : "+color+"\nWeather :\n"+weather, "", "", nil)
}

func main() {

	r := gin.New()

	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(
	// 		http.StatusOK,
	// 		gin.H{"message": "Hello World!"},
	// 	)
	// })

	// r.GET("/books", getbookHandler)

	// r.POST("/books", postbookHandler)

	// r.DELETE("/books/:id", deletebookHandler)

	// r.GET("/color", getcolor)

	r.POST("/noti", noti)

	r.Run()
}
