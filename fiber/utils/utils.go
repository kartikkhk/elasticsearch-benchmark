package utils

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

var client elasticsearch.Client

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
