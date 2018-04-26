/*
Command line stock info retrieval tool by Bryce Harrison
*/

package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	//"net/url"
)
func main() {

	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Please provide a stock symbol")
		os.Exit(1)
	}

	apiUrl := "https://api.iextrading.com/1.0/"
	requestUrl := apiUrl + "stock/" + arguments[1] + "/quote"

	fmt.Println(requestUrl)

	req, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req) 
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	fmt.Println(resp.Body)
}