package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func saveToGoogleSheet(request Strike_Meta_Request_Structure) string {
	userID := request.Bybrisk_session_variables.UserId
	dateCreated := time.Now().Format("2006-01-02 15:04:05")
	userName := request.Bybrisk_session_variables.Username
	userMobile := request.Bybrisk_session_variables.Phone
	teamName := request.User_session_variables.TeamName
	teamEmail := request.User_session_variables.TeamEmail
	teamTheme := request.User_session_variables.TeamTheme
	teamSize := request.User_session_variables.TeamSize
	data := userID + ";" + dateCreated + ";" + userName + ";" + userMobile + ";" + teamName + ";" + teamEmail + ";" + teamTheme[0] + ";" + teamSize[0]
	data = strings.Replace(data, " ", "%20", -1)
	response, err := http.Get(Conf.GSApi1 + "?data=" + data)
	if err != nil {
		fmt.Println("[registrationBot][ERROR]: ", err)
	}

	type googleSheetResponse struct {
		Status         string `json:"status"`
		LastWrittenRow int64  `json:"lastWrittenRow"`
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("[registrationBot][ERROR]: ", err)
	}

	var resp googleSheetResponse
	if err := json.Unmarshal(responseData, &resp); err != nil {
		fmt.Println("[registrationBot][ERROR]: ", err)
	}

	if resp.Status != "OK" {
		return "Failed to register! Try again"
	}
	return "Team " + teamName + " successfully registered!"
}

func saveIndividualMemberToGoogleSheet(request Strike_Meta_Request_Structure, size string) string {

	teamSizeInt, err := strconv.Atoi(size)
	if err != nil {
		fmt.Println("[registrationBot][ERROR]: ", err)
	}

	var data string
	switch teamSizeInt {
	case 1:
		member1Name := request.User_session_variables.NameMem1
		member1Tshirt := request.User_session_variables.TshirtMem1[0]
		data = member1Name + ";" + member1Tshirt
	case 2:
		member1Name := request.User_session_variables.NameMem1
		member2Name := request.User_session_variables.NameMem2
		member1Tshirt := request.User_session_variables.TshirtMem1[0]
		member2Tshirt := request.User_session_variables.TshirtMem2[0]
		data = member1Name + ";" + member1Tshirt + ";" + member2Name + ";" + member2Tshirt
	case 3:
		member1Name := request.User_session_variables.NameMem1
		member2Name := request.User_session_variables.NameMem2
		member3Name := request.User_session_variables.NameMem3
		member1Tshirt := request.User_session_variables.TshirtMem1[0]
		member2Tshirt := request.User_session_variables.TshirtMem2[0]
		member3Tshirt := request.User_session_variables.TshirtMem3[0]
		data = member1Name + ";" + member1Tshirt + ";" + member2Name + ";" + member2Tshirt + ";" + member3Name + ";" + member3Tshirt
	}

	data = strings.Replace(data, " ", "%20", -1)
	response, err := http.Get(Conf.GSApi2 + "?data=" + data)
	if err != nil {
		fmt.Println("[registrationBot][ERROR]: ", err)
	}

	type googleSheetResponse struct {
		Status         string `json:"status"`
		LastWrittenRow int64  `json:"lastWrittenRow"`
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("[registrationBot][ERROR]: ", err)
	}

	var resp googleSheetResponse
	if err := json.Unmarshal(responseData, &resp); err != nil {
		fmt.Println("[registrationBot][ERROR]: ", err)
	}

	if resp.Status != "OK" {
		return "Failed to update team info! Try again"
	}
	return "Team info updated successfully!"
}
