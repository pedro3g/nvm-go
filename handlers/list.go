package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Releases struct {
	Version string
}

func List() []Releases {
	releasesUrl := "https://nodejs.org/download/release/index.json"

	resp, err := http.Get(releasesUrl)

	if err != nil {
		log.Fatal("Error while fetching releases list")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("Error while reading releases list")
	}

	var releases []Releases
	err = json.Unmarshal(body, &releases)

	if err != nil {
		log.Fatal("Error while parsing releases list")
	}

	for _, releases := range releases {
		fmt.Println(releases.Version)
	}

	return releases
}
