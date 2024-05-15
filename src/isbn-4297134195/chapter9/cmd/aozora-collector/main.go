package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"

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

	pat := regexp.MustCompile(`.*/cards/([0-9]+)/card([0-9]+).html$`)

	doc.Find("ol li a").Each(func(n int, elem *goquery.Selection) {
		token := pat.FindStringSubmatch(elem.AttrOr("href", ""))
		if len(token) != 3 {
			return
		}

		pageURL := fmt.Sprintf("https://www.aozora.gr.jp/cards/%s/card%s.html", token[1], token[2])
		_, zipURL := findAuthorAndZIP(pageURL)
		fmt.Println(zipURL)
	})

	return nil, nil
}

func findAuthorAndZIP(siteURL string) (string, string) {
	// log.Println("query", siteURL)
	res, err := http.Get(siteURL)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return "", ""
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println("Unexpected status code:", res.StatusCode)
		return "", ""
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", ""
	}

	author := doc.Find("table[summary= 作家データ ] tr:nth-child(1) td:nth-child(2)").Text()
	zipURL := ""

	doc.Find("table.download a").Each(func(n int, elem *goquery.Selection) {
		href := elem.AttrOr("href", "")
		if strings.HasSuffix(href, ".zip") {
			zipURL = href
		}
	})

	if zipURL == "" {
		return author, ""
	}

	if strings.HasPrefix(zipURL, "http://") || strings.HasPrefix(zipURL, "https://") {
		return author, zipURL
	}

	u, err := url.Parse(siteURL)
	if err != nil {
		return author, ""
	}

	u.Path = path.Join(path.Dir(u.Path), zipURL)

	return author, u.String()
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
