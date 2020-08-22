package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"sync"
)

var once sync.Once
var listUserAgent []string

func init() {
	once.Do(func() {
		listUserAgent = strings.Split(userAgents, "\n")
	})
}

func check(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func GetRandomUserAgent() string {
	randomIndex := rand.Intn(len(listUserAgent))
	return listUserAgent[randomIndex]
}

func Fetch(url string, userAgent string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	check(err)
	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	check(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	return body
}
