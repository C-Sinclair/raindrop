package main

import (
	"fmt"
	"log"
	"raindrop/pkg"

	"github.com/joho/godotenv"
)

func main() {
  fmt.Println("Raindrop 💧")
  // load .env
  err := godotenv.Load()
  if err != nil {
    log.Fatalln(err)
  }
  collections, err := collection.GetCollections()
  if err != nil {
    log.Fatalln(err)
  }
  fmt.Println("Found ", len(collections), " Collections!")
}

