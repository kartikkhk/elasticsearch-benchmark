package utils

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

var client elasticsearch.Client

type CommentDoc struct {
	PostId int32  `json:"postId"`
	Id     int32  `json:"id"`
	Name   string `name:"name"`
	Email  string `email:"email"`
	Body   string `body:"body"`
}

func ConstructClient() {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
		panic(err)
	}
	log.Println(elasticsearch.Version)
	res, err := client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)
}

// func Insert(index string, _type string, data CommentDoc) {
// 	dataJSON, err := json.Marshall(data)
// 	body := string(dataJSON)
// 	ind, err := esclient.Index().
// 		Index(index).
// 		BodyJson(body).
// 		Do(ctx)
// }
