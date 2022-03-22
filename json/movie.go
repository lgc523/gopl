package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int    `json:"released"`
	Region string `json:"region"`
	Type   string `json:"type,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{
			"气象厅的人们", 2022, "netflix-韩国", "烂漫爱情", []string{"韩流进", "李时雨", "陈夏京", "韩气峻", "金善美", "严.."},
		}}
	//[]byte
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("json marshaling failed:%s", err)
	}
	//[{"Title":"气象厅的人们","released":2022,"region":"netflix-韩国","type":"烂漫爱情","Actors":["韩流进","李时雨","陈夏京","韩气峻","金善美","严.."]}]
	fmt.Printf("%s\n", data)

	indent, err := json.MarshalIndent(movies, "", "	")
	if err != nil {
		log.Fatalf("json marshalIndent failed:%s", err)
	}
	/*
		[
		        {
		                "Title": "气象厅的人们",
		                "released": 2022,
		                "region": "netflix-韩国",
		                "type": "烂漫爱情",
		                "Actors": [
		                        "韩流进",
		                        "李时雨",
		                        "陈夏京",
		                        "韩气峻",
		                        "金善美",
		                        "严.."
		                ]
		        }
		]
	*/
	fmt.Printf("%s\n", indent)

	//unmarshal
	var actors []struct{ Actors []string }
	if err := json.Unmarshal(indent, &actors); err != nil {
		log.Fatalf("json unmarshaling failed: %s", err)
	}
	//[{[韩流进 李时雨 陈夏京 韩气峻 金善美 严..]}]
	fmt.Printf("%v\n", actors)
}
