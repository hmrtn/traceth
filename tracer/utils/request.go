package utils

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func RequestTrace(url, data string) []byte {
	client := &http.Client{}
	var reqData = strings.NewReader(data)
	req, err := http.NewRequest("POST", url, reqData)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return bodyText
}
