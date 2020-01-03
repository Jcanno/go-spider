package engine

import (
	"log"
	"spider/fetcher"
)

// SimpleEngine s
type SimpleEngine struct {
}

// Run 处理请求
func (e *SimpleEngine) Run(seeds ...Request) {
	// var requests []Request
	// for _, r := range seeds {
	// 	requests = append(requests, r)
	// }
	// for len(requests) > 0 {
	// 	r := requests[0]
	// 	requests = requests[1:]
	// 	log.Printf("Fetching %s", r.URL)
	// 	// 请求url
	// 	body, err := fetcher.Fetch(r.URL)
	// 	if err != nil {
	// 		log.Printf("Fetcher: error fetching URL %s %v", r.URL, err)
	// 		continue
	// 	}
	// 	// 将结果提交解析
	// 	parseResult := r.ParserFunc(body)
	// 	requests = append(requests, parseResult.Requests...)
	// 	for _, item := range parseResult.Items {
	// 		log.Printf("Got item %v", item)
	// 	}
	// }

	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		/*
		   log.Printf("Fetching %s",r.Url)

		   body, err := fetcher.Fetch(r.Url)

		   if err != nil {
		       log.Printf("Fetcher: error fetching url %s %v", r.Url, err)
		       continue
		   }

		   parseResult := r.ParserFunc(body)
		*/

		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.URL)

	body, err := fetcher.Fetch(r.URL)

	if err != nil {
		log.Printf("Fetcher: error fetching url %s %v", r.URL, err)
		return ParseResult{}, err
	}

	return r.ParserFunc(body), nil
}
