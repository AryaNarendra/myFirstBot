package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Strike_Meta_Request_Structure struct {

	// Bybrisk variable from strike bot
	//
	Bybrisk_session_variables Bybrisk_session_variables_struct `json: "bybrisk_session_variables"`

	// Our own variable from previous API
	//
	User_session_variables User_session_variables_struct `json: "user_session_variables"`
}

type Bybrisk_session_variables_struct struct {

	// User ID on Bybrisk
	//
	UserId string `json:"userId"`

	// Our own business Id in Bybrisk
	//
	BusinessId string `json:"businessId"`

	// Handler Name for the API chain
	//
	Handler string `json:"handler"`

	// Current location of the user
	//
	Location GeoLocation_struct `json:"location"`

	// Username of the user
	//
	Username string `json:"username"`

	// Address of the user
	//
	Address string `json:"address"`

	// Phone number of the user
	//
	Phone string `json:"phone"`
}

type GeoLocation_struct struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type User_session_variables_struct struct {
	TextInput     string             `json:"textInput"`
	LocationInput GeoLocation_struct `json:"locationInput"`
	NumberInput   string             `json:"numberInput"`
	DateInput     []string           `json:"dateInput"`
	Card          []string           `json:"card"`
}

type AppConfig struct {
	Port  string `json:"port"`
	APIEp string `json:"apiep"`
}

var conf *AppConfig

func main() {
	conf = &AppConfig{Port: ":7001", APIEp: ""}
	// Init Routes
	router := gin.Default()
	router.POST("/strike", Getting_started)

	// Start serving the application
	err := router.Run(conf.Port)
	if err != nil {
		log.Fatal("[Initialize] Failed to start server. Error: ", err)
	}
}

func Getting_started(ctx *gin.Context) {
	ctx.JSON(200, "")
}
