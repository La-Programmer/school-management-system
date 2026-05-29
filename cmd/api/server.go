package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello root route"))
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userID := strings.TrimSuffix(path, "/")

		fmt.Println("The ID is:", userID)

		fmt.Println("Query Params:", r.URL.Query())

		queryParams := r.URL.Query()
		sortBy := queryParams.Get("sortBy")
		key := queryParams.Get("key")
		sortOrder := queryParams.Get("sortOrder")

		fmt.Printf("Sort By: %v, Sort Order: %v, Key: %v\n", sortBy, sortOrder, key)

		w.Write([]byte("Hello GET method on teachers route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST method on teachers route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on teachers route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on teachers route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on teachers route"))
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on students route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST method on students route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on students route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on students route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on students route"))
	default:
		w.Write([]byte("Hello on students route"))
	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on execs route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST method on execs route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on execs route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on execs route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on execs route"))
	default:
		w.Write([]byte("Hello on execs route"))
	}
}

func main() {
	port := ":3000"

	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/teachers/", teachersHandler)

	http.HandleFunc("/students", studentsHandler)

	http.HandleFunc("/execs", execsHandler)

	fmt.Println("Server is running on port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalln("Error starting the server: ", err)
	}
}
