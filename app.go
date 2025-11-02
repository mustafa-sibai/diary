package main

import (
	"diary/Routes"
	"net/http"
)

func main() {
	http.HandleFunc("/test", Routes.TestRouter.Handler)
	http.ListenAndServe(":3000", nil)
}
