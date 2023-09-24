package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	router.GET("/v1/albums", getAlbums)
	router.GET("/v1/albums/:id", getAlbumByID)
	router.POST("/v1/albums", postAlbums)
	router.GET("/v1/example", getExample)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func loadExample(output chan<- string) {
	resp, err := http.Get("https://example.com")
	if err != nil {
		output <- fmt.Sprintf("Error reading response body: %v", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		output <- fmt.Sprintf("Error reading response body: %v", err)
		return
	}

	// Send the response body to the output channel
	output <- string(body)
}

func getExample(c *gin.Context) {
	log.Println("Starting getExample")

	output := make(chan string, 1)
	additionalResults := make(chan string, 2)
	var wg sync.WaitGroup
	wg.Add(2)

	go loadExample(output)

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		additionalResults <- "ok"
	}()
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		additionalResults <- "ok2"
	}()

	wg.Wait()
	close(output)
	close(additionalResults)

	var values []string
	for value := range additionalResults {
		values = append(values, value)
	}
	log.Println("Received values:", values)
	// c.JSON(http.StatusOK, gin.H{"results": results})
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(<-output))
}
