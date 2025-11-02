package Routes

import (
	"diary/Router"
	"net/http"
)

var TestRouter = Router.Router{
	GET:  GET,
	POST: POST,
}

func GET(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("This is a GET request response from /test endpoint"))
}

func POST(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("This is a POST request response from /test endpoint"))
}
