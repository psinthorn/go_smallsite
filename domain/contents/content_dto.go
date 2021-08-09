package domain

import (
	"errors"
	"strings"
	"time"
)

type content struct {
	Id          int64     `json: "id" bson: "id"`
	Title       string    `json: "title" bson: "title"`
	SubTitle    string    `json: "sub_title" bson: "sub_title"`
	Content     string    `json: "content" bson: "content"`
	Section     string    `json: "section" bsin: "section"`
	ContentType string    `json: "content_type" bson: "content_type`
	Category    string    `json: "category" bson: "category"`
	Image       string    `json: "image" bson: "image"`
	Tags        string    `json: "tags" bson: "tags"`
	Author      string    `json: "author" bson: "author"`
	Status      string    `json: "status" bson: "status"`
	CreatedAt   time.Time `json: "created_at"`
	UpdatedAt   time.Time `json: "updated_at"`
}

//
// Validate content input
//

func (ct *content) Validate() error {
	ct.Title = strings.TrimSpace(ct.Title)
	if ct.Title == "" {
		return errors.New("Title can not be empty!")
	}
	return nil
}
