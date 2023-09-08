package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fungsi TestPingRoute untuk menguji rute "/ping"
func TestPingRoute(t *testing.T) {
	// Memanggil fungsi setupRouter untuk membuat router
	router := setupRouter()

	// Membuat objek recorder HTTP baru
	w := httptest.NewRecorder()

	// Membuat HTTP request baru dengan method GET dan target rute "/ping"
	req, _ := http.NewRequest("GET", "/ping", nil)

	// Melakukan serve HTTP dengan router dan recorder HTTP
	router.ServeHTTP(w, req)

	// Memeriksa apakah kode respon HTTP sama dengan 200
	assert.Equal(t, 200, w.Code)

	// Memeriksa apakah isi respon HTTP sama dengan `{"message":"pong"}`
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}
