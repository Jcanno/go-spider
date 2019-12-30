package engine

import (
	"log"
	"spider/fetcher"
)

// Run 处理请求
func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.URL)
		// 请求url
		body, err := fetcher.Fetch(r.URL)
		if err != nil {
			log.Printf("Fetcher: error fetching URL %s %v", r.URL, err)
			continue
		}
		// 将结果提交解析
		parseResult := r.ParserFunc(body)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}
