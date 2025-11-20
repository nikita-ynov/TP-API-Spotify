package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"spotify/pages"
)

type ApiData struct {
	Access_token string `json:"access_token"`
}

type ApiDataAlbums struct {
	Albums struct {
		Items []struct {
			Name         string `json:"name"`
			Release_date string `json:"release_date"`
			Total_tracks int    `json:"total_tracks"`
			Images       []struct {
				Url string `json:"url"`
			}
		} `json:"items"`
	} `json:"albums"`
}

type ApiDataTrack struct {
	Album struct {
		Images []struct {
			Url string `json:"url"`
		}
		Name         string `json:"name"`
		Release_date string `json:"release_date"`
	} `json:"album"`
	Artists []struct {
		Name string `json:"name"`
	} `json:"artists"`
	External_urls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Name string `json:"name"`
}

func renderPage(w http.ResponseWriter, filename string, data any) {
	err := pages.Temp.ExecuteTemplate(w, filename, data)
	if err != nil {
		http.Error(w, "Erreur rendu template : "+err.Error(), http.StatusInternalServerError)
	}
}

func getToken() string {
	apiUrl := "https://accounts.spotify.com/api/token?grant_type=client_credentials&client_id=cf4022930b814fb88b78a2f7836e779c&client_secret=9b1bd4aa9ec84d9c8ced17acf6bf3fcd"
	resp, err := http.Post(apiUrl, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, errBody := io.ReadAll(resp.Body)
	if errBody != nil {
		fmt.Println("Error: ", errBody.Error())
	}

	var decodeData ApiData
	json.Unmarshal(body, &decodeData)

	return "Bearer  " + decodeData.Access_token
}

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(getToken())
	renderPage(w, "index.html", nil)
}

func AlbumDamso(w http.ResponseWriter, r *http.Request) {
	apiUrl := "https://api.spotify.com/v1/search/?q=Damso&type=album"

	token := getToken()

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Add("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	var decodeData ApiDataAlbums
	json.Unmarshal(body, &decodeData)

	// fmt.Printf("Albums: %+v\n", decodeData.Albums.Items)
	data := map[string]any{
		"Albums": decodeData.Albums.Items,
	}
	renderPage(w, "albumDamso.html", data)
}

func TrackLaylow(w http.ResponseWriter, r *http.Request) {
	apiUrl := "https://api.spotify.com/v1/tracks/67Pf31pl0PfjBfUmvYNDCL"
	token := getToken()

	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Add("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		return
	}

	var decodeData ApiDataTrack
	json.Unmarshal(body, &decodeData)

	data := map[string]any{
		"Tracks": decodeData,
	}
	// fmt.Println(decodeData)
	renderPage(w, "trackLaylow.html", data)
}
