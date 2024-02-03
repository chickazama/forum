package main

import (
	"database/sql"
	"log"
	"matthewhope/forum/api"
	"matthewhope/forum/repo"
	"matthewhope/forum/sqlite"
	"matthewhope/forum/ui"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./sqlite/data/Identity.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	businessDb, err := sql.Open("sqlite3", "./sqlite/data/Business.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = sqlite.PostsTableDown(businessDb)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = sqlite.UsersTableUp(db)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = sqlite.PostsTableUp(businessDb)
	if err != nil {
		log.Fatal(err.Error())
	}
	mux := http.NewServeMux()
	serveStaticFiles(mux)
	mux.Handle("/", ui.NewIndexHandler())
	mux.Handle("/users", api.NewUsersHandler(repo.NewSQLiteRepository()))
	err = http.ListenAndServe(":7777", mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func serveStaticFiles(mux *http.ServeMux) {
	fsRoot := http.Dir("./static/")
	fs := http.FileServer(fsRoot)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}
