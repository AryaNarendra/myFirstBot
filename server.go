package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/strike-official/go-sdk/strike"
)

type Strike_Meta_Request_Structure struct {

	// Bybrisk variable from strike bot
	//
	Bybrisk_session_variables Bybrisk_session_variables_struct `json: "bybrisk_session_variables,omitempty"`

	// Our own variable from previous API
	//
	User_session_variables User_session_variables_struct `json: "user_session_variables,omitempty"`
}

type Bybrisk_session_variables_struct struct {

	// User ID on Bybrisk
	//
	UserId string `json:"userId,omitempty"`

	// Our own business Id in Bybrisk
	//
	BusinessId string `json:"businessId,omitempty"`

	// Handler Name for the API chain
	//
	Handler string `json:"handler,omitempty"`

	// Current location of the user
	//
	Location GeoLocation_struct `json:"location,omitempty"`

	// Username of the user
	//
	Username string `json:"username,omitempty"`

	// Address of the user
	//
	Address string `json:"address,omitempty"`

	// Phone number of the user
	//
	Phone string `json:"phone,omitempty"`
}

type GeoLocation_struct struct {
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

type User_session_variables_struct struct {
	TeamName            string   `json:"teamName,omitempty"`
	TeamEmail           string   `json:"teamEmail,omitempty"`
	TeamTheme           []string `json:"teamTheme,omitempty"`
	TeamSize            []string `json:"teamSize,omitempty"`
	SaveDetailsResponse []string `json:"save_details_response,omitempty"`
	NameMem1            string   `json:"name_mem1,omitempty"`
	NameMem2            string   `json:"name_mem2,omitempty"`
	NameMem3            string   `json:"name_mem3,omitempty"`
	TshirtMem1          []string `json:"tshirt_mem1,omitempty"`
	TshirtMem2          []string `json:"tshirt_mem2,omitempty"`
	TshirtMem3          []string `json:"tshirt_mem3,omitempty"`
}

var Conf *AppConfig

func main() {
	err := InitAppConfig("config.json")
	if err != nil {
		log.Fatal("[registrationBot][ERROR]: ", err)
	}
	Conf = GetAppConfig()
	// Init Routes
	router := gin.Default()
	router.POST("/registrationBot", Getting_started)
	router.POST("/registrationBot/saveDetails", saveDetails)
	router.POST("/registrationBot/saveTeamDetails", saveTeamDetails)

	// Start serving the application
	err = router.Run(Conf.Port)
	if err != nil {
		log.Fatal("[Initialize] Failed to start server. Error: ", err)
	}
}

func Getting_started(ctx *gin.Context) {
	var request Strike_Meta_Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("[registrationBot][ERROR]: ", err)
	}
	fmt.Println("[registrationBot][/]: ", request)
	name := request.Bybrisk_session_variables.Username
	// Core Logic
	strikeObj := strike.Create("getting_started", Conf.ApiEp+"/saveDetails")

	quesObj := strikeObj.Question("teamName").
		QuestionText().
		SetTextToQuestion("Hi "+name+"! provide your team name", "desc")

	quesObj.Answer(true).TextInput("")

	quesObj2 := strikeObj.Question("teamEmail").
		QuestionText().
		SetTextToQuestion("Enter your team email", "desc")

	quesObj2.Answer(true).TextInput("")

	quesObj3 := strikeObj.Question("teamTheme").
		QuestionText().
		SetTextToQuestion("Choose a theme/idea for the competition", "desc")

	quesObj3.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Healthcare", "#3b5375", true).AddTextRowToAnswer(strike.H5, "Healthcare units nearby based on your health insurance", "#4d4d4d", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Ecommerce", "#3b5375", true).AddTextRowToAnswer(strike.H5, "GTT for e-commerce (e-commerce price notifier)", "#4d4d4d", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Marketplace", "#3b5375", true).AddTextRowToAnswer(strike.H5, "A personal e-commerce bot", "#4d4d4d", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Social Media", "#3b5375", true).AddTextRowToAnswer(strike.H5, "Fact checker bot", "#4d4d4d", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Personal Assistance", "#3b5375", true).AddTextRowToAnswer(strike.H5, "A point of contact for you. Automates daily errands", "#4d4d4d", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Entertainment", "#3b5375", true).AddTextRowToAnswer(strike.H5, "Movie recommendation bot", "#4d4d4d", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Social", "#3b5375", true).AddTextRowToAnswer(strike.H5, "Job recommendation bot", "#4d4d4d", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Game", "#3b5375", true).AddTextRowToAnswer(strike.H5, "Quiz bot", "#4d4d4d", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Finance", "#3b5375", true).AddTextRowToAnswer(strike.H5, "Crypto price predictor (or Mutual fund recommendation for the ones playing safe)", "#4d4d4d", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Chat server/client", "#3b5375", true).AddTextRowToAnswer(strike.H5, "Chat service", "#4d4d4d", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Polling Bot", "#3b5375", true).AddTextRowToAnswer(strike.H5, "A bot which can take a poll and provide result to the poll creator", "#4d4d4d", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Meme Autobot", "#3b5375", true).AddTextRowToAnswer(strike.H5, "A bot which can create meme out of any image Tourism", "#4d4d4d", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Tourism", "#3b5375", true).AddTextRowToAnswer(strike.H5, "A smart location based itinerary builder", "#4d4d4d", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Wikipedia", "#3b5375", true).AddTextRowToAnswer(strike.H5, "A Wikipedia bot which stops a particular Article from vandalism", "#4d4d4d", false).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Other", "#3b5375", true)
	quesObj4 := strikeObj.Question("teamSize").
		QuestionText().
		SetTextToQuestion("Enter your team size", "desc")

	quesObj4.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "1", "#3b5375", true).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "2", "#3b5375", true).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "3", "#3b5375", true)

	ctx.JSON(200, strikeObj)
}

func saveDetails(ctx *gin.Context) {
	var request Strike_Meta_Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("[registrationBot][ERROR]: ", err)
	}
	fmt.Println("[registrationBot][/ -> /saveDetails]: ", request)

	resp := saveToGoogleSheet(request)
	strikeObj := strike.Create("save_details", Conf.ApiEp+"/saveTeamDetails?size="+request.User_session_variables.TeamSize[0])
	quesObj := strikeObj.Question("save_details_response").QuestionText().
		SetTextToQuestion(resp, "desc")
	quesObj.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION).
		AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Enter Team details", "#3b5375", true)

	teamSizeInt, err := strconv.Atoi(request.User_session_variables.TeamSize[0])
	if err != nil {
		fmt.Println("[registrationBot][ERROR]: ", err)
	}

	for i := 0; i < teamSizeInt; i++ {
		counter := strconv.Itoa(i + 1)
		quesObj1 := strikeObj.Question("name_mem"+counter).QuestionText().
			SetTextToQuestion("Name of member "+counter+" ?", "desc")

		quesObj1.Answer(true).TextInput("")

		quesObj2 := strikeObj.Question("tshirt_mem"+counter).QuestionText().
			SetTextToQuestion("T-shirt size for member "+counter+" ?", "desc")

		quesObj2.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION).
			AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "S", "#3b5375", true).
			AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "M", "#3b5375", true).
			AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "L", "#3b5375", true).
			AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "XL", "#3b5375", true).
			AnswerCard().SetHeaderToAnswer(1, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "XXL", "#3b5375", true)
	}

	ctx.JSON(200, strikeObj)
}

func saveTeamDetails(ctx *gin.Context) {
	var request Strike_Meta_Request_Structure
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("[registrationBot][ERROR]: ", err)
	}
	fmt.Println("[registrationBot][/saveDetails -> /saveTeamDetails]: ", request)

	teamSize := ctx.Query("size")
	resp := saveIndividualMemberToGoogleSheet(request, teamSize)

	strikeObj := strike.Create("save_team_details", "")
	quesObj := strikeObj.Question("").QuestionText().
		SetTextToQuestion(resp, "desc")

	quesObj.Answer(false).AnswerCardArray(strike.VERTICAL_ORIENTATION).
		AnswerCard().SetHeaderToAnswer(10, strike.HALF_WIDTH).AddGraphicRowToAnswer(strike.PICTURE_ROW, []string{Conf.HackathonMemePicURL}, []string{}).
		AnswerCard().SetHeaderToAnswer(10, strike.HALF_WIDTH).AddTextRowToAnswer(strike.H4, "Read more about the hackathon on \n\nhttps://strike.bybrisk.com/hackathon", "#3b5375", false)

	ctx.JSON(200, strikeObj)
}
