package server

import (
	"fmt"
	// "io"
	"net/http"
	// "net/http/httptest"
)

func TestingServer() {
	// var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	// 	// w.Header().Set("Content-Type", "text/plain")
	// 	// w.Write([]byte("Hello World"))
	// 	fmt.Fprint(w, "Hello World!")
	// }

	// server := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: handler,
	// }

	// error := server.ListenAndServe()
	// if error != nil {
	// 	panic(error)
	// }
}

func TestingServeMux() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.Method, " ")
		fmt.Fprintln(w, r.URL.Path)
	})
	mux.HandleFunc("/hi/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hai, Halo!")
	})
	mux.HandleFunc("/hi/hihi/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "2xHi!")
	})
	mux.HandleFunc("/hi/hihi/hihihi/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "3xHi!")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	error := server.ListenAndServe()
	if error != nil {
		panic(error)
	}
}

func HandlerTesting() {
	// var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello World!")
	// }

	// request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	// recorder := httptest.NewRecorder()

	// handler(recorder, request)

	// response := recorder.Result()
	// body, _ := io.ReadAll(response.Body)

	// bodyString := string(body)
	// fmt.Println(bodyString)
}