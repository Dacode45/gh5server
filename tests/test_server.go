package main
import (
  "net/http/httptest"
  "net/http"
)
func TestMyHandler(t *testing.T) {
  	router := NewRouter()
    server := httptest.NewServer(&router)
    defer server.Close()

    resp, err := http.Post("http://localhost:8080/courts")

}
