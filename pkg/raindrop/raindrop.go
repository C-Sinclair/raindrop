package raindrop

import (
	"encoding/json"
	"fmt"
	"log"
	"raindrop/pkg/collection"
	"raindrop/pkg/request"
)

type raindropsRes struct {
	Result bool
	Items  []Raindrop
}

type raindropRes struct {
	Result bool
	Item   Raindrop
}

type Raindrop struct {
	Id         int `json:"_id"` // The id of the Collection
	Collection RaindropCollection

	Cover      string // Cover URL
	Created    string // when the raindrop was created
	Domain     string // hostname of a link. Files always have raindrop.io hostname
	Excerpt    string // description
	LastUpdate string // When the raindrop was updated
	Link       string // url
	Media      []RaindropMedia
	Tags       []string // tags list
	Title      string   // Name of the collection
	Type       string   // link article image video document or audio
	User       collection.User
}

type RaindropCollection struct {
	// id of the collection the raindrop resides in
	Id int `json:"$id"`
}

type RaindropMedia struct {
	Link string // url of cover
}

func GetRaindrops(search string) ([]Raindrop, error) {
	fmt.Println("Getting raindrops...")

	query := "?perpage=50"
	if len(search) > 0 {
		query += fmt.Sprintf("&search=%s", search)
	}
	url := fmt.Sprintf("/raindrops/0%s", query)
	page := 1
	var results []Raindrop
	for {
		res, err := getPaginatedRaindrops(url, page)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, res.Items...)
		if res.Result == true && len(res.Items) > 0 {
			page = page + 1
			continue
		}
		break
	}
	return results, nil
}

func getPaginatedRaindrops(url string, page int) (*raindropsRes, error) {
	get_url := fmt.Sprintf("%s&page=%d", url, page)
	// fmt.Println("Getting from ", get_url)
	res, err := request.GetRequest(get_url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	drops := new(raindropsRes)
	err = json.NewDecoder(res.Body).Decode(drops)

	return drops, err
}

func GetRaindrop(id string) (Raindrop, error) {
	url := fmt.Sprintf("/raindrop/%s", id)
	res, err := request.GetRequest(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	raindrop := new(raindropRes)
	err = json.NewDecoder(res.Body).Decode(raindrop)

	return raindrop.Item, err
}
