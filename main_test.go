package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAlbums(t *testing.T) {
	// Create a new Gin router for testing
	router := gin.Default()
	router.GET("/v1/albums", getAlbums)

	// Create a request to the /v1/albums endpoint
	req, _ := http.NewRequest("GET", "/v1/albums", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the status code and response body
	assert.Equal(t, http.StatusOK, w.Code)

	var response []album
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// You can add more specific assertions based on your expected response
}

func TestPostAlbums(t *testing.T) {
	// Create a new Gin router for testing
	router := gin.Default()
	router.POST("/v1/albums", postAlbums)

	// Create a new album to post
	newAlbum := album{
		ID:     "4",
		Title:  "New Album",
		Artist: "Test Artist",
		Price:  29.99,
	}
	// Convert the album to JSON
	payload, err := json.Marshal(newAlbum)
	assert.NoError(t, err)

	// Create a request to the /v1/albums endpoint with the JSON payload
	req, _ := http.NewRequest("POST", "/v1/albums", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the status code and response body
	assert.Equal(t, http.StatusCreated, w.Code)

	var response album
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// You can add more specific assertions based on your expected response
}

func TestGetAlbumByID(t *testing.T) {
	// Create a new Gin router for testing
	router := gin.Default()
	router.GET("/v1/albums/:id", getAlbumByID)

	// Define a test album ID
	testAlbumID := "1"

	// Create a request to the /v1/albums/:id endpoint with the test album ID
	req, _ := http.NewRequest("GET", "/v1/albums/"+testAlbumID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the status code and response body
	assert.Equal(t, http.StatusOK, w.Code)

	var response album
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// You can add more specific assertions based on your expected response
}

func TestGetAlbumByIDNotFound(t *testing.T) {
	// Create a new Gin router for testing
	router := gin.Default()
	router.GET("/v1/albums/:id", getAlbumByID)

	// Define a non-existent test album ID
	testAlbumID := "nonexistent"

	// Create a request to the /v1/albums/:id endpoint with the test album ID
	req, _ := http.NewRequest("GET", "/v1/albums/"+testAlbumID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert the status code and response body
	assert.Equal(t, http.StatusNotFound, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// You can add more specific assertions based on your expected response
}
