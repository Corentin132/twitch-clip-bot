package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/twitch"
)

var (
	clientID     = "my_client_id"
	clientSecret = "My_secret"
	oauth2Config *clientcredentials.Config
)

type Video struct {
	Id            string
	Stream_id     string
	User_id       string
	User_login    string
	User_name     string
	Tittle        string
	Description   string
	Created_at    string
	Published_at  string
	Url           string
	Thumbnail_url string
	Viewable      string
}
type AllVideo struct {
	Data []Video
}
type Categorie struct {
	Box_art_url string //`json:"box_bart_url"`
	Id          string //`json:"id"`
	Name        string //`json:"name"`
}
type Search struct {
	Data []Categorie //`json:"data"`
}
type Channel struct {
	Brocaster_language string
	Brocaster_login    string
	Display_name       string
	Game_id            string
	Game_name          string
	Id                 string
	Is_live            bool
	Tags_ids           []string
	Thumbnail_url      string
	Title              string
	Started_at         string
}
type AllChannels struct {
	Data []Channel
}

func getChannelID(accessToken string, username string, target AllChannels) string {
	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/search/channels?query="+username, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Client-Id", "my_token")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		log.Fatalln("error  getting Channel ID : ", err)
	}
	fmt.Println("get vidéo of : ", resp.Body)
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("error  getting category : ", err)
	}
	err = json.Unmarshal(bodyText, &target)
	if err != nil {
		panic(err)
	}
	return target.Data[0].Id

}
func getCategoryID(accessToken string, target Search) string {

	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/search/categories?query=justchatting", nil)
	if err != nil {
		// handle err
		log.Fatalln("error  getting category : ", err)

	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Client-Id", "my_token")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		log.Fatalln("error  getting category : ", err)
	}
	defer resp.Body.Close()

	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("error  getting category : ", err)
	}
	err = json.Unmarshal(bodyText, &target)
	if err != nil {
		panic(err)
	}
	return target.Data[0].Id
}
func getVideoChannel(accessToken string, id string, target AllVideo, date string) AllVideo {
	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/clips?broadcaster_id="+id+"&language=fr&first=5&sort=views"+"&started_at="+date+"T00:00:00Z", nil)
	if err != nil {
		// handle err
		log.Fatalln("error getting video ")
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Client-Id", "my_token")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		log.Fatalln("error getting video ")
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("error getting video ")
	}
	err = json.Unmarshal(bodyText, &target)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("All info about clips getting ✅")
	return target
}
func getVideoGame(accessToken string, id string, target AllVideo) AllVideo {
	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/clips?game_id="+id+"&language=fr"+"&first=5&started_at=2023-01-03T00:00:00Z", nil)
	if err != nil {
		// handle err
		log.Fatalln("error getting video ")
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Client-Id", "my_token")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
		log.Fatalln("error getting video ")
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("error getting video ")
	}
	err = json.Unmarshal(bodyText, &target)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("All info about clips getting ✅")
	return target
}

func getAccessToken() string {
	oauth2Config = &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     twitch.Endpoint.TokenURL,
	}
	token, err := oauth2Config.Token(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Access token: %s\n", token.AccessToken)
	return token.AccessToken
}
func main() {

	var research Search
	var allVideo AllVideo
	var AllChannels AllChannels
	accessToken := getAccessToken()

	id := getCategoryID(accessToken, research)
	username := ""
	date := ""
	fmt.Printf("Streamer name : ")
	_, err := fmt.Scanf("%s", &username)
	if err != nil {
		log.Fatalln("Error getting username scanf", err)
	}
	fmt.Printf("Clip date (ex: 2023-06-27) : ")
	_, err = fmt.Scanf("%s", &date)
	if err != nil {
		log.Fatalln("Error getting date scanf", err)
	}

	chanelID := getChannelID(accessToken, username, AllChannels)

	fmt.Println("chanel ID: ", chanelID)
	fmt.Println("id:", id)

	allVideo = getVideoChannel(accessToken, chanelID, allVideo, date)
	// dataLength := len(allVideo.Data)
	for i, clips := range allVideo.Data {
		// dowload all clips
		fmt.Println("Title : ", clips.Tittle)
		fmt.Println("Author : ", clips.User_name)
		fmt.Println("Url : ", clips.Url)
		DowloadVideo(clips.Url)

		//----------------------------------------------- for wait the shell script finished ----------------------------
		cmd := exec.Command("./urlScript.sh")
		stdout, _ := cmd.Output()
		fmt.Println(string(stdout))
		//------------------------------------------------for wait the shell script finished------------------------------

		fmt.Printf("Dowloaded %d/%d \n", i+1, len(allVideo.Data))

	}
	// getNewTitle("alderiate", "？! LA GAME LA PLUS LONGUE V2 ？! [1456750651]")
	log.Println("THE END ! 👋")

}
