package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func HandleRequest(w *httptest.ResponseRecorder, r *http.Request) { //storage Storage
	router := getRouter() // storage
	router.ServeHTTP(w, r)
}

func createTestAlbum() album {
	testAlbum := album{ID: "2", Title: "Blue Train", Artist: "John Coltraine", Price: 56.99}
	storage.Create(testAlbum)
	return testAlbum
}

func TestAlbumsList(t *testing.T) {
	request, _ := http.NewRequest("GET", "/albums", strings.NewReader(""))
	w := httptest.NewRecorder()
	HandleRequest(w, request) //storage
	if w.Code != http.StatusOK {
		t.Fatal("status not okg")
	}
}

func TestAlbumsDetail(t *testing.T) {
	testAlbum := createTestAlbum()
	request, _ := http.NewRequest("GET", "/albums/"+testAlbum.ID, strings.NewReader(""))
	w := httptest.NewRecorder()
	HandleRequest(w, request) //storage
	if w.Code != http.StatusOK {
		t.Fatal("status not okg")
	}
}

func TestAlbumsNotFound(t *testing.T) {
	albumId := "9999"
	request, _ := http.NewRequest("GET", "/albums/"+albumId, strings.NewReader(""))
	w := httptest.NewRecorder()
	HandleRequest(w, request) //storage
	if w.Code != http.StatusNotFound {
		t.Fatal("status must be 404")
	}
}

func TestDeleteAlbums(t *testing.T) {
	testAlbum := createTestAlbum()
	request, _ := http.NewRequest("DELETE", "/albums/"+testAlbum.ID, strings.NewReader(""))
	w := httptest.NewRecorder()
	HandleRequest(w, request) //storage
	if w.Code != http.StatusNoContent {
		t.Fatal("status must be 204")
	}
}

func TestDeleteNotFound(t *testing.T) {
	albumId := "9999"
	request, _ := http.NewRequest("DELETE", "/albums/"+albumId, strings.NewReader(""))
	w := httptest.NewRecorder()
	HandleRequest(w, request) //storage
	if w.Code != http.StatusNotFound {
		t.Fatal("status must be 404")
	}
}

func TestUpdateAlbumsNotFound(t *testing.T) {
	albumId := "999"
	request, _ := http.NewRequest("PUT", "/albums/"+albumId, strings.NewReader(`{"title": "test"}`))
	w := httptest.NewRecorder()
	HandleRequest(w, request) //storage
	if w.Code != http.StatusNotFound {
		t.Fatal("status must be 404", w.Code)
	}
}

func TestUpdateAlbums(t *testing.T) {
	testAlbum := createTestAlbum()
	request, _ := http.NewRequest("PUT", "/albums/"+testAlbum.ID, strings.NewReader(`{"title": "test"}`))
	w := httptest.NewRecorder()
	HandleRequest(w, request) //storage
	if w.Code != http.StatusOK {
		t.Fatal("status must be ok", w.Code)
	}
}

func TestCreateBadStructure(t *testing.T) {
	request, _ := http.NewRequest("POST", "/albums", strings.NewReader(""))
	w := httptest.NewRecorder()
	HandleRequest(w, request) //storage
	if w.Code != http.StatusBadRequest {
		t.Fatal("status must be 400", w.Code)
	}
}

func TestCreateAlbums(t *testing.T) {
	request, _ := http.NewRequest("POST", "/albums", strings.NewReader(`{"id": "4", "title": "The Modern Sound of Betty Carter", "artist": "Betty Carter", "price": 39.99}`))
	w := httptest.NewRecorder()
	HandleRequest(w, request) //storage

	if w.Code != http.StatusCreated {
		t.Fatal("status must be 201", w.Code)
	}
}
