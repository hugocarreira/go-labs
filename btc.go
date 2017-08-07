package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Dados struct {
	price   string
	success bool
	time    string
}

func main() {

	response, erro := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	if erro != nil {
		fmt.Printf("Request with error%s\n", erro)
	} else {
		responseJson, _ := ioutil.ReadAll(response.Body)

		fmt.Println(string(responseJson))

		data := Dados{}
		erro = json.Unmarshal(responseJson, &data)

		if erro != nil {
			fmt.Println(erro)
		}
		fmt.Println(data)
	}

}
