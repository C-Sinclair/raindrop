package collection

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const BASE_URL = "https://api.raindrop.io/rest/v1"

type collectionsRes struct {
  result bool
  items []Collection
}

type Collection struct {
  _id int // The id of the Collection
  access Access
  collaborators Collaborators // When this object is present, means that collections is shared. Content of this object is private and not very useful.
  color string // primary color of the collection as HEX  
  count int // count of raindrops in collection
  cover []string // Collection cover URL. This array always have one item due to legacy reasons
  created string // when the collection is created
  expanded bool // Whether the collection’s sub-collections are expanded
  lastUpdate string // When the collection is updated
  parent Parent
  public bool // Collection and raindrops that it contains will be accessible without authentication by public link
  sort int // The order of collection (descending). Defines the position of the collection among all the collections with the same parent.$id
  title string // Name of the collection
  user User 
  /** 
  view style of Collection
  - list (default)
  - simple
  - grid
  - masonry (Pinterest like grid)
  */
  view string 
}

type Access struct {
  /**
  1 - read only access (equal to public=true)
  2 - collaborator with read only access
  3 - collaborator with write only access
  4 - owner
  */
  level int
  draggable bool // Does it possible to change parent of this collection?
}

type Collaborators struct {}

type Parent struct {
  // the id of the parent collection. Not specified for root collections 
  id int `json:"$id"`
}

type User struct {
  // owner id 
  id int `json:"$id"`
}

func GetCollections() {
  print("Get my collections")

  url := fmt.Sprintf("%s/collections", BASE_URL)
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
  defer res.Body.Close()

  // TODO: get collections decoding into the struct
  collections := new(collectionsRes)
  json.NewDecoder(res.Body).Decode(collections)

  str, _ := json.Marshal(collections)
  log.Printf(string(str))
}
