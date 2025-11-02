package Routes

import (
	"diary/Models"
	"diary/Router"
	"diary/Types"
	"encoding/json"
	"net/http"
)

var FetchAllPostsRouter = Router.Router{
	GET: fetchAllPosts,
}

func fetchAllPosts(res http.ResponseWriter, req *http.Request) {

	var request Types.FetchAllPostsRequest

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid request payload"))
		return
	}

	posts, err := Models.FetchAllPosts(request)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Failed to fetch posts: " + err.Error()))
		return
	}

	response := Types.FetchAllPostsResponse{Posts: posts}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(response)
}
