package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	_ClientCompany := false
	var ClientID string = ""
	var Link string = "https://api.crm.solarz.com.br/api/v2/open-api/client/TYPE/ID"
	strings.ReplaceAll(Link, "ID", ClientID)
	if _ClientCompany {
		strings.ReplaceAll(Link, "TYPE", "organizations")
	}else{
		strings.ReplaceAll(Link, "TYPE", "persons")
	}
	HttpResult, Error := http.Get(Link)
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
		BodyLines := strings.Split(Body,",")
	}
}
