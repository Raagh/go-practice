package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"quiet/hn"
	"sort"
	"strings"
	"time"
)

func main() {
	// parse flags
	var port, numStories int
	flag.IntVar(&port, "port", 3000, "the port to start the web server on")
	flag.IntVar(&numStories, "num_stories", 30, "the number of top stories to display")
	flag.Parse()

	tpl := template.Must(template.ParseFiles("./index.gohtml"))

	http.HandleFunc("/", handler(numStories, tpl))

	// Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func getTopStories(numStories int) ([]item, error) {
	var client hn.Client
	ids, err := client.TopItems()
	if err != nil {
		return nil, errors.New("Failed to load top stories")
	}

	var stories []item
	stories = getStories(ids[0:numStories])

	at := 0
	for len(stories) < numStories {
		need := (numStories - len(stories)) * 5 / 4 // this adds a few extra stories just in case in the 30 we get discussions
		stories = append(stories, getStories(ids[at:at+need])...)
		at += need
	}

	return stories, nil
}

func getStories(ids []int) []item {
	var client hn.Client
	type result struct {
		index int
		item  item
		err   error
	}

	resultCh := make(chan result)

	for i := 0; i < len(ids); i++ {
		go func(index, id int) {
			hnItem, err := client.GetItem(id)
			if err != nil {
				resultCh <- result{index: index, err: err}
			}

			item := parseHNItem(hnItem)
			resultCh <- result{index: index, item: item}
		}(i, ids[i])
	}

	var results []result
	for i := 0; i < len(ids); i++ {
		results = append(results, <-resultCh)
	}

	sort.Slice(results, func(i int, j int) bool {
		return results[i].index < results[j].index
	})

	var stories []item
	for i := 0; i < len(ids); i++ {
		result := results[i]
		if result.err != nil {
			continue
		}

		if isStoryLink(result.item) {
			stories = append(stories, result.item)
			if len(stories) >= len(ids) {
				break
			}
		}
	}

	return stories
}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		stories, err := getTopStories(numStories)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		data := templateData{
			Stories: stories,
			Time:    time.Now().Sub(start),
		}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to process the template", http.StatusInternalServerError)
			return
		}
	})
}

func isStoryLink(item item) bool {
	return item.Type == "story" && item.URL != ""
}

func parseHNItem(hnItem hn.Item) item {
	ret := item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

// item is the same as the hn.Item, but adds the Host field
type item struct {
	hn.Item
	Host string
}

type templateData struct {
	Stories []item
	Time    time.Duration
}
