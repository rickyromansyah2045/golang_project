package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func api() {

	url := "https://dev-api.unipin.com/in-game-topup/list"
	method := "POST"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("partnerid", "9b42a14d-a986-40a9-b4cc-354be6aea6db")
	req.Header.Add("timestamp", "1566552295")
	req.Header.Add("path", "in-game-topup/list")
	req.Header.Add("auth", "1d9f5e7aca9f3c14da7c957d6977447739877cebfc10fcf3682bd32da47a2bda")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
