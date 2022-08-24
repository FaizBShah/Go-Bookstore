package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/FaizBShah/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func loadEnv() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(filepath.Join(dir, ".env"))

	if err != nil {
		panic(err)
	}
}

func main() {
	loadEnv()

	router := mux.NewRouter()
	routes.RegisterBookStore(router)

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
