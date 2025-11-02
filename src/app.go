package main

import (
	"diary/Database"
	"diary/Routes"
	"log"
	"net/http"
)

func main() {
	err := Database.InitializeDatabase()
	if err != nil {
		log.Fatal(err)
	}

	defer Database.GetSqlDB().Close()

	http.HandleFunc("/test", Routes.TestRouter.Handler)
	http.HandleFunc("/add-new-post", Routes.AddNewPostRouter.Handler)
	http.HandleFunc("/fetch-all-posts", Routes.FetchAllPostsRouter.Handler)
	http.ListenAndServe(":3000", nil)
}
