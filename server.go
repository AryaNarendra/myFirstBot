package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/strike-official/go-sdk/strike"
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
	TeamName  string   `json:"teamName"`
	TeamEmail string   `json:"teamEmail"`
	TeamTheme []string `json:"teamTheme"`
	TeamSize  string   `json:"teamSize"`
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
	router.POST("/", Getting_started)
	router.POST("/saveDetails", saveDetails)

	// Start serving the application
	err := router.Run(conf.Port)
	if err != nil {
		log.Fatal("[Initialize] Failed to start server. Error: ", err)
	}
}

func Getting_started(ctx *gin.Context) {
	var request Strike_Meta_Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Err")
	}
	name := request.Bybrisk_session_variables.Username
	// Core Logic
	strikeObj := strike.Create("getting_started", "http://3556-2405-201-a407-908e-a123-da20-e212-3472.ngrok.io/saveDetails")

	quesObj := strikeObj.Question("teamName").
		QuestionText().
		SetTextToQuestion("Hi! "+name+" Provide your Team Name", "desc")

	quesObj.Answer(true).TextInput("")

	quesObj2 := strikeObj.Question("teamEmail").
		QuestionText().
		SetTextToQuestion("Enter your Team Email", "desc")

	quesObj2.Answer(true).TextInput("")

	quesObj3 := strikeObj.Question("teamTheme").
		QuestionText().
		SetTextToQuestion("Choose a Theme/Idea", "desc")

	quesObj3.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Theme1", "black", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Theme2", "black", true).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H5, "Theme3", "black", false)

	quesObj4 := strikeObj.Question("teamSize").
		QuestionText().
		SetTextToQuestion("Enter your Team Size", "desc")

	quesObj4.Answer(true).NumberInput("")

	ctx.JSON(200, strikeObj)
}

func saveDetails(ctx *gin.Context) {
	var request Strike_Meta_Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Err")
	}
	fmt.Println(request.User_session_variables.TeamName)
	fmt.Println(request.User_session_variables.TeamEmail)

}
