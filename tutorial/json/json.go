package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	marshalling()
	unmarshalling()
}

/*
	MARSHALLING
*/
func marshalling() {
	type gResponse struct {
		ResponseData struct {
			Results []struct {
				GsearchResultClass string `json:"GsearchResultClass"`
				UnescapedURL       string `json:"unescapedUrl"`
				URL                string `json:"url"`
				VisibleURL         string `json:"visibleUrl"`
				CacheURL           string `json:"cacheUrl"`
				Title              string `json:"title"`
				TitleNoFormatting  string `json:"titleNoFormatting"`
				Content            string `json:"content"`
			} `json:"results"`
		} `json:"responseData"`
	}

	var a = struct {
		GsearchResultClass string `json:"GsearchResultClass"`
		UnescapedURL       string `json:"unescapedUrl"`
		URL                string `json:"url"`
		VisibleURL         string `json:"visibleUrl"`
		CacheURL           string `json:"cacheUrl"`
		Title              string `json:"title"`
		TitleNoFormatting  string `json:"titleNoFormatting"`
		Content            string `json:"content"`
	}{
		GsearchResultClass: "GwebSearch",
		UnescapedURL:       "https://www.reddit.com/r/golang",
		URL:                "https://www.reddit.com/r/golang",
		VisibleURL:         "www.reddit.com",
		CacheURL:           "http://www.google.com/search?q=cache:W...",
		Title:              "r/\u003cb\u003eGolang\u003c/b\u003e - Reddit",
		TitleNoFormatting:  "r/Golang - Reddit",
		Content:            "First Open Source",
	}

	r := gResponse{
		ResponseData: struct {
			Results []struct {
				GsearchResultClass string `json:"GsearchResultClass"`
				UnescapedURL       string `json:"unescapedUrl"`
				URL                string `json:"url"`
				VisibleURL         string `json:"visibleUrl"`
				CacheURL           string `json:"cacheUrl"`
				Title              string `json:"title"`
				TitleNoFormatting  string `json:"titleNoFormatting"`
				Content            string `json:"content"`
			} `json:"results"`
		}{
			Results: []struct {
				GsearchResultClass string `json:"GsearchResultClass"`
				UnescapedURL       string `json:"unescapedUrl"`
				URL                string `json:"url"`
				VisibleURL         string `json:"visibleUrl"`
				CacheURL           string `json:"cacheUrl"`
				Title              string `json:"title"`
				TitleNoFormatting  string `json:"titleNoFormatting"`
				Content            string `json:"content"`
			}{
				{
					// ????
				},
			},
		},
	}

	barr, err := json.Marshal(r)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s\n", string(barr))
}

/*
	UNMARSHALLING
*/
func unmarshalling() {
	type Result struct {
		ResponseData struct {
			Results []struct {
				GsearchResultClass string `json:"GsearchResultClass"`
				UnescapedURL       string `json:"unescapedUrl"`
				URL                string `json:"url"`
				VisibleURL         string `json:"visibleUrl"`
				CacheURL           string `json:"cacheUrl"`
				Title              string `json:"title"`
				TitleNoFormatting  string `json:"titleNoFormatting"`
				Content            string `json:"content"`
			} `json:"results"`
		} `json:"responseData"`
	}

	byt := []byte(`
		{
			"responseData":{
				"results":[
						{
							"GsearchResultClass":"GwebSearch",
							"unescapedUrl":"https://www.reddit.com/r/golang",
							"url":"https://www.reddit.com/r/golang",
							"visibleUrl":"www.reddit.com",
							"cacheUrl":"http://www.google.com/search?q=cache:W...",
							"title":"r/\u003cb\u003eGolang\u003c/b\u003e - Reddit",
							"titleNoFormatting":"r/Golang - Reddit",
							"content":"First Open Source"
						},
						{
							"GsearchResultClass":"GwebSearch",
							"unescapedUrl":"http://tour.golang.org/",
							"url":"http://tour.golang.org/",
							"visibleUrl":"tour.golang.org",
							"cacheUrl":"http://www.google.com/search?q=cache:O...",
							"title":"A Tour of Go",
							"titleNoFormatting":"A Tour of Go",
							"content":"Welcome to a tour of the Go programming ..."
						}
				]
			}
		}
`)
	var s Result
	fmt.Printf("%+v\n", s)

	err := json.Unmarshal(byt, &s)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v\n", s)
}
