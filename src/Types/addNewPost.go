package Types

type AddNewPostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type AddNewPostResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
