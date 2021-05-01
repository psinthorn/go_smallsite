package contents

type Content struct {
	Id          int64  `json: "id" bson: "id"`
	Title       string `json: "title" bson: "title"`
	SubTitle    string `json: "sub_title" bson: "sub_title"`
	Content     string `json: "content" bson: "content"`
	ContentType string `json: "content_type" bson: "content_type`
	Category    string `json: "category" bson: "category"`
	Image       string `json: "image" bson: "image"`
	Tags        string `json: "tags" bson: "tags"`
	Author      string `json: "author" bson: "author"`
	Status      string `json: "status" bson: "status"`
	Created_At  string `json: "created_at" bson: "created_at"`
	Updated_At  string `json: "updated_at" bson: "updated_at"`
}

//
// Validate content input
//

// func (content *Content) Validate() *errors.ErrorRespond {
// 	content.Title = strings.TrimSpace(content.Title)
// 	if content.Title == "" {
// 		return errors.NewContentAlertNotice("Title can not be empty!")
// 	}
// 	return nil
// }
