package main

import (
	"spider/engine"
	"spider/parser"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	engine.Run(engine.Request{
		URL:        url,
		ParserFunc: parser.ParseCityList,
	})
}
