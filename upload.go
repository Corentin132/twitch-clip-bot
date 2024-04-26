package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Choices struct {
	text          string
	index         int
	logprobs      string
	finish_reason string
}
type GTPResponse struct {
	id      string
	object  string
	created int
	model   string
	choices []Choices
	usage   struct {
		prompt_token     int
		completion_token int
		total_token      int
	}
}

func getNewTitle(broadcaster string, title string) {
	API_KEY := "my_api_key"
	body := strings.NewReader(`{
		"model": "text-davinci-003",
		"prompt": "écrit moi un titre accrocheur pour une vidéo sur le streameur ` + broadcaster + ` en référence au nom du fichier de la vidéo "` + title + `" en enlevant les chiffres de la fin",
		"temperature": 0,
		"max_tokens": 256,
	  }`)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", body)
	if err != nil {
		log.Println("getting title error: ", err)
	}
	req.Header.Set("Authorization", "Bearer "+API_KEY)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("getting title error: ", err)
	}
	log.Printf("Response ==== %s", resp.Body)
	defer resp.Body.Close()

	// manque plus qu'a clean la data avec les structs déjà fait

}
func changeTitle() {
	files, err := ioutil.ReadDir("./clips")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())

		// apelle api icii ----------------------------------------------------------------
		New_Path := "gfg.txt"
		e := os.Rename(file.Name(), New_Path)
		if e != nil {
			log.Fatal(e)
		}
	}
}
func uploadVideo() {
	// CLIENT_KEY := "my_key"
	// CLIENT_SECRET := "my_secret"
	// CODE := "Twitch1234"

	// Rename and Remove a file
	// Using Rename() function

}
