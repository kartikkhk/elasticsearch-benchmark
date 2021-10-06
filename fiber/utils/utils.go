package utils

import (
	"fmt"

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

func ConstructClient() {
	_, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))
	if err != nil {
		panic(err)
	}
	fmt.Println("elastic client connected")
}

// func Insert(index string, _type string, data CommentDoc) {
// 	dataJSON, err := json.Marshall(data)
// 	body := string(dataJSON)
// 	ind, err := esclient.Index().
// 		Index(index).
// 		BodyJson(body).
// 		Do(ctx)
// }
