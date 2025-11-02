package Models

import (
	"diary/Database"
	"diary/Schema"
	"diary/Types"
)

func FetchAllPosts(request Types.FetchAllPostsRequest) ([]Types.Post, error) {
	return Schema.FetchAllPosts(Database.GetSqlDB())
}
