package Routes

import (
	"diary/Models"
	"diary/Router"
	"diary/Types"
	"encoding/json"
	"net/http"
)

var AddNewPostRouter = Router.Router{
	POST: addNewPost,
}

func addNewPost(res http.ResponseWriter, req *http.Request) {

	var request Types.AddNewPostRequest

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid request payload"))
		return
	}

	if err := Models.AddNewPost(request); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Failed to add new post"))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte("New post added successfully"))
}
