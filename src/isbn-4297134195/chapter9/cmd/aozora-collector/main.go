package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Entry struct {
	AuthorID string
	Author  string
	TitleID string
	Title string
	InfoURL string
	ZipURL string
}

func findEntries(siteURL string) ([]Entry, error) {
	res, err := http.Get(siteURL)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Unexpected status code:", res.StatusCode)
		return nil, fmt.Errorf("Unexpected status code: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("ol li a").Each(func(n int, elem *goquery.Selection) {
		fmt.Println(elem.Text(), elem.AttrOr("href", ""))
	})

	return nil, nil
}

func main() {
	listURL := "https://www.aozora.gr.jp/index_pages/person879.html"

	entries, err := findEntries(listURL)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		fmt.Println(entry.Title, entry.ZipURL)
	}
}
