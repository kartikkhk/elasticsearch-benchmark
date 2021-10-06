package utils

import (
	"github.com/olivere/elastic/v7"
)

var client elastic.Client

type CommentDoc struct {
	PostId int32  `json:"postId"`
	Id     int32  `json:"id"`
	Name   string `name:"name"`
	Email  string `email:"email"`
	Body   string `body:"body"`
}

func GetESClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	return client, err
}
