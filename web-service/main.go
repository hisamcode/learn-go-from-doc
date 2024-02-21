package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// album represent data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// album slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Traing", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaighan", Price: 39.99},
}

// toJson create json from any and returns []byte and error
func toJson(data interface{}) ([]byte, error) {
	b, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return nil, err
	}
	return b, nil
}

// receiveJson receive JSON
func receiveJson(rc io.ReadCloser, data interface{}) error {
	err := json.NewDecoder(rc).Decode(data)
	if err != nil {
		return err
	}
	return nil
}

// getAlbums responds with the list of all album as json
func getAlbums(w http.ResponseWriter, r *http.Request) {
	b, err := toJson(albums)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(w http.ResponseWriter, r *http.Request) {
	var newAlbum album
	receiveJson(r.Body, &newAlbum)
	albums = append(albums, newAlbum)

	b, err := toJson(newAlbum)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// getAlbumByID locates the album whose ID value matches the id
// parHelloameter sent by the client, then returns that album as a response
func getAlbumByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	for _, a := range albums {
		if a.ID == id {
			b, err := toJson(a)
			if err != nil {
				log.Println(err)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(b)
			return
		}
	}
	b, err := toJson(map[string]string{
		"message": "album not found",
	})
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	w.Write(b)

}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /albums", getAlbums)
	mux.HandleFunc("GET /albums/{id}", getAlbumByID)
	mux.HandleFunc("POST /albums", postAlbums)

	fmt.Println("Listen on 8080")
	err := http.ListenAndServe("127.0.0.1:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
