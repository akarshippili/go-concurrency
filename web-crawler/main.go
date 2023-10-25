package main

import (
	"log"
	"sync"

	"github.com/akarshippili/go-concurrency/web-crawler/fetcher"
)

type Crawler struct {
	visited   map[string]bool
	n         int
	readCount int
	mu        sync.Mutex
	wg        sync.WaitGroup
}

func (crawler *Crawler) crawl(url string) {

	procede := true
	defer crawler.wg.Done()

	// critical section - cs
	crawler.mu.Lock()
	count := crawler.readCount
	_, ok := crawler.visited[url]
	if count < crawler.n && !ok {
		crawler.visited[url] = true
		crawler.readCount += 1
	} else {
		procede = false
	}
	crawler.mu.Unlock()

	if procede {
		log.Default().Printf("fetching url: [%v]\n", url)
		relatedLinks := fetcher.GetRelatedLinks(url)
		crawler.wg.Add(len(relatedLinks))
		log.Default().Printf("related url: [%v]", relatedLinks)

		for _, relatedLink := range relatedLinks {
			go crawler.crawl(relatedLink)
		}
	} else {
		var reason string
		if count >= crawler.n {
			reason = "reached limit"
		} else if ok {
			reason = "revisted"
		} else {
			reason = "unknow"
		}
		log.Default().Printf("Exiting - %v\n", reason)
	}
}

func main() {
	// content, err := fetcher.GetBody("http://amazon.com/")
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }

	// relatedUrls := fetcher.GetRelatedLinks("http://amazon.com/")

	// log.Default().Println(content)
	// log.Default().Println(relatedUrls)

	crawler := Crawler{
		visited:   make(map[string]bool),
		n:         1000,
		readCount: 0,
	}

	crawler.crawl("http://amazon.com/")
	crawler.wg.Wait()
}
