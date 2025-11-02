package Router

import "net/http"

type Router struct {
	GET  func(http.ResponseWriter, *http.Request)
	POST func(http.ResponseWriter, *http.Request)
}

func (router *Router) Handler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		if router.GET == nil {
			http.Error(res, "Not Found", http.StatusNotFound)
			return
		}
		router.GET(res, req)
	case "POST":
		if router.POST == nil {
			http.Error(res, "Not Found", http.StatusNotFound)
			return
		}
		router.POST(res, req)
	default:
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
