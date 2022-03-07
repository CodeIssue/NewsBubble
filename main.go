package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	response, err := http.Get("https://www.youtube.com/channel/UCsBjURrPoezykLs9EqgamOA")
	if err != nil {
		log.Print(err)
	}
	responseBody, _ := ioutil.ReadAll(response.Body)

	log.Print(response.Status)
	log.Print(string(responseBody))

}
