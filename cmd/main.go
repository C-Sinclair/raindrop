package main

import (
	"log"
	"raindrop/pkg"

	"github.com/joho/godotenv"
)

func main() {
  print("Raindrop 💧")
  // load .env
  err := godotenv.Load()
  if err != nil {
    log.Fatalln(err)
  }
  collection.GetCollections()
}

