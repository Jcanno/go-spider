package parser

import (
	"regexp"
	"spider/engine"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

// ParseCity 解析城市信息
func ParseCity(contents []byte) engine.ParseResult {

	re := regexp.MustCompile(cityRe)
	all := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, c := range all {
		result.Items = append(result.Items, "User:"+string(c[2])) //用户名字

		result.Requests = append(result.Requests, engine.Request{
			URL:        string(c[1]),
			ParserFunc: engine.NilParser,
		})
	}

	return result
}
