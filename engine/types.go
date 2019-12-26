package engine

// ParseResult 解析后返回的结果
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

// Request 请求
type Request struct {
	URL        string                   // 解析出来的URL
	ParserFunc func([]byte) ParseResult // 处理这个URL所需要的函数
}

// NilParser 是
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
