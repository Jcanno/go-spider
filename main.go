package main

import (
	"spider/engine"
	"spider/parser"
	"spider/scheduler"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		URL:        url,
		ParserFunc: parser.ParseCityList,
	})
}
