package main

import (
	"fmt"
	"location/helpers"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env variables")
	}

	engine := gin.New()

	engine.LoadHTMLGlob("templates/*.html")

	engine.Use(gin.Logger())

	engine.GET("/", indexHandler)
	engine.GET("/location", DisplayLocation)

	engine.Run(":3000")
}

func DisplayLocation(c *gin.Context) {

	longitude := c.Query("long")
	latitude := c.Query("lat")

	if latitude == "" || longitude == "" {
		c.String(http.StatusBadRequest, "Latitude and longitude are required.")
		return
	}

	//mapbox token
	accessToken := os.Getenv("TOKEN")

	// Base URL for the Mapbox Geocoding API
	baseURL := "https://api.mapbox.com/geocoding/v5/mapbox.places/"

	//creating url object
	u, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		c.String(http.StatusInternalServerError, "Error parsing URL")
		return
	}

	// Create query parameters using url.Values
	queryParams := url.Values{}
	queryParams.Add("access_token", accessToken)

	// Add latitude and longitude as part of the path
	u.Path = u.Path + longitude + "," + latitude + ".json"

	// Set the query parameters for the URL
	u.RawQuery = queryParams.Encode()

	// Convert the URL object back to a string
	finalURL := u.String()

	//call helper and get result
	response, err := helpers.SendRequestAndFindPlace(finalURL)
	if err != nil {
		c.String(http.StatusInternalServerError, "error in server")
		return
	}

	fmt.Println("Final URL:", finalURL)

	c.HTML(http.StatusOK, "location.html", response)
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

type Data struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
