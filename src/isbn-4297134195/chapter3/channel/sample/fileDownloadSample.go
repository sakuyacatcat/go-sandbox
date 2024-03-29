package sample

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func downloadCSV(wg *sync.WaitGroup, urls []string, ch chan []byte) {
	defer wg.Done()
	defer close(ch)

	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			log.Println("could not download CSV:", err)
			continue
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			resp.Body.Close()
			log.Println("could not read content:", err)
			continue
		}
		resp.Body.Close()
		ch <- b
	}
}

func main() {
	urls := []string{
		"https://example.com/1.csv",
		"https://example.com/2.csv",
		"https://example.com/3.csv",
		"https://example.com/4.csv",
		"https://example.com/5.csv",
		"https://example.com/6.csv",
		"https://example.com/7.csv",
		"https://example.com/8.csv",
		"https://example.com/9.csv",
		"https://example.com/10.csv",
	}

	ch := make(chan []byte)

	var wg sync.WaitGroup

	wg.Add(1)
	go downloadCSV(&wg, urls, ch)

	for b := range ch {
		r := csv.NewReader(bytes.NewReader(b))
		for {
			records, err := r.Read()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(records)
		}
	}
	wg.Wait()
}
