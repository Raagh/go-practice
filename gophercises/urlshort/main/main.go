package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"urlshort"

	_ "github.com/lib/pq"
)

func main() {
	// yamlFile := flag.String("yamlFile", "", "Provide a yaml file to read the http handler paths")
	// jsonFile := flag.String("jsonFile", "", "Provide a json file to read the http handler paths")
	flag.Parse()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	db, err := sql.Open("postgres", "postgres://your_username:your_username@host:port/your_database_name?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to PostgreSQL!")

	var id int
	err = db.QueryRow("INSERT INTO mappings(path, url) VALUES($1, $2) RETURNING id", "/urlshort-godoc", "https://godoc.org/github.com/gophercises/urlshort").Scan(&id)
	if err != nil {
		panic(err)
	}

	dbHandler, err := urlshort.DBHandler(db, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", dbHandler)

	// if *yamlFile != "" {
	// 	file, err := os.ReadFile(*yamlFile)
	// 	if err != nil {
	// 		fmt.Printf("It was no possible to open the file: %s\n", *yamlFile)
	// 		os.Exit(1)
	// 	}
	//
	// 	yamlHandler, err := urlshort.YAMLHandler(file, mapHandler)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Starting the server on :8080")
	// 	http.ListenAndServe(":8080", yamlHandler)
	// } else if *jsonFile != "" {
	// 	file, err := os.ReadFile(*jsonFile)
	// 	if err != nil {
	// 		fmt.Printf("It was no possible to open the file: %s\n", *jsonFile)
	// 		os.Exit(1)
	// 	}
	//
	// 	jsonHandler, err := urlshort.JSONHandler(file, mapHandler)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("Starting the server on :8080")
	// 	http.ListenAndServe(":8080", jsonHandler)
	// } else {
	// 	fmt.Println("Starting the server on :8080")
	// 	http.ListenAndServe(":8080", mapHandler)
	// }
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
