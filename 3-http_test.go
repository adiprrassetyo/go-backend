package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func TestHttp() {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
