package Models

import (
	"diary/Database"
	"diary/Schema"
	"diary/Types"
)

func AddNewPost(request Types.AddNewPostRequest) error {
	return Schema.AddNewPost(Database.GetSqlDB(), request.Title, request.Content)
}
