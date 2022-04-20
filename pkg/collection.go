package raindrop

import (
	"encoding/json"
	"fmt"
	"log"
)

type collectionsRes struct {
	Result bool
	Items  []Collection
}

type Collection struct {
	Id            int `json:"_id"` // The id of the Collection
	Access        Access
	Author        *bool          // user is the author
	Collaborators *Collaborators // When this object is present, means that collections is shared. Content of this object is private and not very useful.
	Color         *string        // primary color of the collection as HEX
	Count         int            // count of raindrops in collection
	Cover         []string       // Collection cover URL. This array always have one item due to legacy reasons
	Created       string         // when the collection is created
	// creatorRef
	Expanded   bool // Whether the collection’s sub-collections are expanded
	LastAction string
	LastUpdate string // When the collection is updated
	Parent     *Parent
	Public     bool // Collection and raindrops that it contains will be accessible without authentication by public link
	Slug       string
	Sort       int    // The order of collection (descending). Defines the position of the collection among all the collections with the same parent.$id
	Title      string // Name of the collection
	User       User
	/**
	  view style of Collection
	  - list (default)
	  - simple
	  - grid
	  - masonry (Pinterest like grid)
	*/
	View string
}

type Access struct {
	/**
	  1 - read only access (equal to public=true)
	  2 - collaborator with read only access
	  3 - collaborator with write only access
	  4 - owner
	*/
	Level     int
	Draggable bool // Does it possible to change parent of this collection?
	Root      *bool
}

type Collaborators struct{}

type Parent struct {
	// the id of the parent collection. Not specified for root collections
	Id int `json:"$id"`
}

type User struct {
	// owner id
	Id int `json:"$id"`
}

func GetCollections() ([]Collection, error) {
	fmt.Println("Getting collections...")

	res, err := GetRequest("/collections")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	collections := new(collectionsRes)
	err = json.NewDecoder(res.Body).Decode(collections)

	return collections.Items, err
}
