package engine

import (
	"log"
	"spider/fetcher"
)

// Scheduler 接口
type Scheduler interface {
	Submit(request Request)
	ConfigureMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

// ConcurrentEngine ConcurrentEngine
type ConcurrentEngine struct {
	Scheduler   Scheduler //Sheduler
	WorkerCount int       //worker的数量
}

// Run Run
// func (e *ConcurrentEngine) Run(seeds ...Request) {

// 	//worker公用一个in，out
// 	in := make(chan Request)
// 	out := make(chan ParseResult)

// 	e.Scheduler.ConfigureMasterWorkerChan(in)

// 	for i := 0; i < e.WorkerCount; i++ {
// 		createWorker(in, out) //创建worker
// 	}

// 	//参数seeds的request，要分配任务
// 	for _, r := range seeds {
// 		e.Scheduler.Submit(r)
// 	}

// 	itemCount := 0
// 	//从out中获取result，对于item就打印即可，对于request，就继续分配
// 	for {
// 		result := <-out
// 		for _, item := range result.Items {
// 			log.Printf("Got %d  item : %v", itemCount, item)
// 			itemCount++
// 		}

// 		for _, request := range result.Requests {
// 			e.Scheduler.Submit(request)
// 		}
// 	}
// }

// Run Run
func (e *ConcurrentEngine) Run(seeds ...Request) {

	//worker公用一个in，out
	//in := make(chan Request)
	out := make(chan ParseResult)

	//e.Scheduler.ConfigureMasterWorkerChan(in)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		//createWorker(in, out) //创建worker
		createWorker(out, e.Scheduler)
	}

	//参数seeds的request，要分配任务
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	//从out中获取result，对于item就打印即可，对于request，就继续分配
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got %d  item : %v", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

// createWorker 创建worker
// func createWorker(in chan Request, out chan ParseResult) {
// 	go func() {
// 		for {
// 			request := <-in
// 			result, err := worker(request)
// 			if err != nil {
// 				continue
// 			}
// 			out <- result
// 		}
// 	}()
// }

// createWorker 创建worker
func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			//需要让scheduler知道已经就绪了
			s.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
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
