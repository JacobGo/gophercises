package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"

	"./hn"
)

func main() {
	// parse flags
	var port, numStories int
	flag.IntVar(&port, "port", 3000, "the port to start the web server on")
	flag.IntVar(&numStories, "num_stories", 30, "the number of top stories to display")
	flag.Parse()

	fmt.Printf("Starting server on port http://localhost:%d with %d stories...\n", port, numStories)
	tpl := template.Must(template.ParseFiles("./index.gohtml"))

	http.HandleFunc("/", handler(numStories, tpl))

	// Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

var (
	ids []int
	topErr error
	cachedTopItemsExpiration time.Time
	cachedStories map[int]hn.Item
	cachedStoriesMutex = sync.RWMutex{}
)
func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var client hn.Client

		if start.After(cachedTopItemsExpiration) || cachedTopItemsExpiration.IsZero() {
			ids, topErr = client.TopItems()
			cachedTopItemsExpiration = start.Add(time.Minute * 15)
		}
		if topErr != nil {
			http.Error(w, "Failed to load top stories", http.StatusInternalServerError)
			return
		}

		if cachedStories == nil {
			cachedStories = make(map[int]hn.Item)
		}

		var stories []item
		type result struct {
			item item
			idx int
			err error
		}
		var validResults []result
		results := make(chan result)
		batch := 0
		for len(validResults) < numStories {
			newBatch := batch + (numStories * 5/4)
			for i := batch; i < newBatch; i++ {
				go func(i int) {
					var hnItem hn.Item
					if cachedItem, ok := cachedStories[ids[i]]; ok {
						hnItem = cachedItem
					} else {
						newItem, err := client.GetItem(ids[i])
						if err != nil {
							results <- result{idx: i, err: err}
							return
						} else {
							cachedStoriesMutex.Lock()
							cachedStories[ids[i]] = newItem
							hnItem = newItem
							cachedStoriesMutex.Unlock()
						}
					}

					item := parseHNItem(hnItem)
					results <- result{idx: i, item: item}
				}(i)
			}
			for i := batch; i < newBatch; i++ {
				result := <- results
				if result.err == nil && isStoryLink(result.item) {
					validResults = append(validResults, result)
				}
			}
			batch = newBatch
		}


		sort.Slice(validResults, func(i,j int) bool {
			return validResults[i].idx < validResults[j].idx
		})

		for i := 0; i < numStories; i++ {
			stories = append(stories, validResults[i].item)
		}


		data := templateData{
			Stories: stories,
			Time:    time.Now().Sub(start),
		}
		err := tpl.Execute(w, data)
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