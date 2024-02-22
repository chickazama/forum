package main

import (
	"fmt"
	"log"
	"matthewhope/forum/api"
	"matthewhope/forum/repo"
	"matthewhope/forum/ui"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// db, err := sql.Open("sqlite3", "./sqlite/data/Identity.db")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// businessDb, err := sql.Open("sqlite3", "./sqlite/data/Business.db")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// err = sqlite.PostsTableDown(businessDb)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// err = sqlite.UsersTableUp(db)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// err = sqlite.PostsTableUp(businessDb)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// p := transport.PostDTO(models.Post{PostID: 1, Category: "Example", Content: "Test"})
	// fmt.Println(p)
	mux := http.NewServeMux()
	serveStaticFiles(mux)
	r := repo.NewSQLiteRepository()
	mux.Handle("/", ui.NewIndexHandler(r))
	mux.Handle("/posts", api.NewPostsHandler(r))
	mux.Handle("/users", api.NewUsersHandler(r))
	fmt.Println("Setting up server...")
	err := http.ListenAndServe(":7777", mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func serveStaticFiles(mux *http.ServeMux) {
	fsRoot := http.Dir("./static/")
	fs := http.FileServer(fsRoot)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
}
