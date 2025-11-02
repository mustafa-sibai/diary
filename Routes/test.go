package Routes

import (
	"diary/Router"
	"fmt"
	"net/http"
)

var TestRouter = Router.Router{
	GET: GET,
}

func GET(res http.ResponseWriter, req *http.Request) {
	fmt.Println("GET request received")
}
