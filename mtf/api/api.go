package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type JSONResponse struct {
	Results []Movie
}

type Movie struct {
	Overview string
	Title    string
}

func Search(title string, apiKey string, language string) (titles []string) {
	fmt.Printf("Searching %q ...\n", title)
	title = strings.ReplaceAll(title, " ", "+")
	url := "https://api.themoviedb.org/3/search/movie?api_key=" + apiKey + "&language=" + language + "&query=" + title
	log.Println("querying " + url)
	var searchResult JSONResponse

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println("response status: " + resp.Status)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(body, &searchResult); err != nil {
		log.Fatal(err)
	}

	if len(searchResult.Results) != 0 {
		log.Println("Found:")
		for i := 0; i < len(searchResult.Results) && i < 5; i++ {
			log.Printf("	%q\n", searchResult.Results[i].Title)
			titles = append(titles, searchResult.Results[i].Title)

		}
	}
	return
}
