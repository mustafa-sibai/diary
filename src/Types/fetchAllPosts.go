package Types

type Post struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type FetchAllPostsRequest struct {
	// No fields needed for this request
}

type FetchAllPostsResponse struct {
	Posts []Post `json:"posts"`
}
