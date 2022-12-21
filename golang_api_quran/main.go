package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	response, err := http.Get("https://api.npoint.io/99c279bb173a6e28359c/data")

	if err != nil {
		fmt.Println("Error Get Data", err)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		println("Error", err)
	}

	fmt.Println(string(responseData))

}
