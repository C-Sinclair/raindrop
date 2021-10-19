package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"raindrop/pkg/raindrop"

	"github.com/joho/godotenv"
)

/**
TODO: Add CLI flags to filter by Collection
*/
func main() {
  fmt.Println("Raindrop 💧")
  // load .env
  err := godotenv.Load()
  if err != nil {
    log.Fatal(err)
  }
  // user raindrops search
  raindrops, err := raindrop.GetRaindrops("")
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Found ", len(raindrops), " matching that search")
  selected := doSearch(func(in io.WriteCloser) {
    for _, drop := range raindrops {
      line := fmt.Sprintf("%d => %s\t-- %s\t -- %s", drop.Id, drop.Title, drop.Excerpt, drop.Tags)
      fmt.Fprintln(in, line)
      time.Sleep(5 * time.Millisecond)
    }
  })
  id := strings.Split(selected, " ")[0]
  raindrop, err := raindrop.GetRaindrop(id)
  url := raindrop.Link
  err = exec.Command("xdg-open", url).Start()
  if err != nil {
    log.Fatal(err)
  }
}

func doSearch(input func(in io.WriteCloser)) string {
  shell := os.Getenv("SHELL")
  if len(shell) == 0 {
    shell = "sh"
  }
  cmd := exec.Command(shell, "-c", "fzf --layout reverse")
  cmd.Stderr = os.Stderr
  in, _ := cmd.StdinPipe()
  go func() {
    input(in)
    in.Close()
  }()
  result, _ := cmd.Output()
  return string(result)
}
