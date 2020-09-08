package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Basic Unit tests
func Test_getBooks(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/books", nil)
	w := httptest.NewRecorder()
	getBooks(w, r)
	//assert
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)
}

func Test_getBook(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/books/1", nil)
	w := httptest.NewRecorder()

	getBook(w, r)
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)
}

func Test_updateBooks(t *testing.T) {
	r, _ := http.NewRequest("PUT", "/api/books/1", nil)
	w := httptest.NewRecorder()
	updateBooks(w, r)
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 5, w.Body.Len())
}

func Test_deleteBook(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/api/books/2", nil)
	w := httptest.NewRecorder()
	deleteBook(w, r)
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 5, w.Body.Len())
}
func Test_deleteBooks(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/api/books", nil)
	w := httptest.NewRecorder()
	deleteBooks(w, r)
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 5, w.Body.Len())
}

func Test_getBookTitles(t *testing.T) {
	r, _ := http.NewRequest("GET", "/api/books/titles/", nil)
	w := httptest.NewRecorder()
	getBookTitles(w, r)
	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 15, w.Body.Len())
}
