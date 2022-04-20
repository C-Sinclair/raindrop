package raindrop

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const BASE_URL = "https://api.raindrop.io/rest/v1"

func GetRequest(path string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", BASE_URL, path)
	token := fmt.Sprintf("Bearer %s", os.Getenv("ACCESS_TOKEN"))
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	return res, err
}
