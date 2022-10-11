package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://dev-api.unipin.com/in-game-topup/list"
	method := "POST"

	payload := strings.NewReader(``)

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)

	partnerid := "9b42a14d-a986-40a9-b4cc-354be6aea6db"

	timestamp := "1566552295"

	path := "in-game-topup/list"

	secretKey := "w56kbwxuxh3heka3"

	a := hmac.New(sha256.New, []byte(secretKey))
	a.Write([]byte(partnerid))

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("partnerid", partnerid)
	req.Header.Add("timestamp", timestamp)
	req.Header.Add("path", path)
	req.Header.Add("auth", auth)
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
