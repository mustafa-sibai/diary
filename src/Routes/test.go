package Routes

import (
	"diary/Router"
	"net/http"
)

var TestRouter = Router.Router{
	GET:  test_get,
	POST: test_post,
}

func test_get(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("This is a GET request response from /test endpoint"))
}

func test_post(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("This is a POST request response from /test endpoint"))
}
