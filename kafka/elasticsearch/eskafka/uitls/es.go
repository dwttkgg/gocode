package main

import (
	"context"
	"fmt"
	elastic "github.com/olivere/elastic/v7"
)

// Elasticsearch demo

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func initEs() (err error) {
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.30.102:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("connect to es success")
	p1 := Person{Name: "lmh", Age: 18, Married: false}
	for i := 1; i < 10000; i++ {
		put1, err := client.Index().
			Index("user").
			Type("user").
			Id(fmt.Sprintf("%d", i)).
			BodyJson(p1).
			Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	}
	return
}
