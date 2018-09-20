package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const topAPI = "https://hacker-news.firebaseio.com/v0/topstories.json"
const itemAPI = "https://hacker-news.firebaseio.com/v0/item/%s.json"

func main() {
	resp, err := http.Get(topAPI)

	if err != nil {
		log.Fatal(err)
	}

	var s string

	b, err := ioutil.ReadAll(resp.Body)
	s = string(b)

	// Using the json package might work better
	s = strings.Replace(s, "[", "", 1)
	s = strings.Replace(s, "]", "", 1)

	var topstories = strings.Split(s, ",")

	for i := range topstories {
		resp, err := http.Get(fmt.Sprintf(itemAPI, topstories[i]))
		if err != nil {
			log.Fatal(err)
		}
		b, err := ioutil.ReadAll(resp.Body)

		file, err := os.Create(fmt.Sprintf("%s.json", topstories[i]))
		if err != nil {
			log.Fatal("Cannot create file", err)
		}
		defer file.Close()

		fmt.Fprintf(file, string(b))

	}
}
