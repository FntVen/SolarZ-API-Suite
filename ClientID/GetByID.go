package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// I Hate It reformats my code after save and i cant find a way to turn it off
func main() {
	var INIT_Date string = ""
	var END_Date string = ""
	var Limit string = ""
	var FID string = ""
	var _ScopeStage bool = false
	//var StageID string = ""
	var Link string = "https://api.crm.solarz.com.br/api/v2/open-api/deal?startDate=SDate&endDate=TodayDate&page=0&size=Quantidade&sort=%5B%22string%22%5D&pipelineId=ID"
	strings.ReplaceAll(Link, "SDate", INIT_Date)
	strings.ReplaceAll(Link, "TodayDate", END_Date)
	strings.ReplaceAll(Link, "Quantidade", Limit)
	strings.ReplaceAll(Link, "ID", FID)
	HttpResult, Error := http.Get("")
	if Error != nil {
		fmt.Println("Error of type ", Error)
	}
	defer HttpResult.Body.Close()
	Body, Error := io.ReadAll(HttpResult.Body)
	if Error != nil {
		fmt.Println("Error of type ", Error)
	}
	if !_ScopeStage {
		fmt.Println(string(Body))
	} else {
		//Implement splitting lines and filtering
	}
}
