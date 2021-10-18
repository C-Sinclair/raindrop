package main

import (
	"encoding/json"
	"fmt"
	"log"

	// "raindrop/pkg/collection"
	"raindrop/pkg/raindrop"

	"github.com/joho/godotenv"
)

func main() {
  // fmt.Println("Raindrop 💧")
  // load .env
  err := godotenv.Load()
  if err != nil {
    log.Fatal(err)
  }
  // get user collections
  // collections, err := collection.GetCollections()
  // if err != nil {
  //   log.Fatal(err)
  // }
  // fmt.Println("Found ", len(collections), " Collections!")
  // user raindrops search
  raindrops, err := raindrop.GetRaindrops("")
  if err != nil {
    log.Fatal(err)
  }
  // fmt.Println("Found ", len(raindrops), " matching that search")
  res, err := json.Marshal(raindrops)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(res))
}

