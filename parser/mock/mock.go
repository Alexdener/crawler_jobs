package mock

import (
	"fmt"

	"crawler_jobs/engine"
	"crawler_jobs/parser"
)

//
//  ParseMockRequest
//  @Description: yapi mock data (docker-compose => https://github.com/fjc0k/docker-YApi)
//  @param bytes
//  @param s
//  @return engine.ParseResult
//
func ParseMockRequest(bytes []byte, s string) engine.ParseResult {
	url := string(bytes)
	requests := make([]engine.Request, 0)
	requests = append(requests, engine.Request{
		Url:    url,
		Parser: parser.NewFuncParser(ParseMockRequest, "ParseMockRequest"),
	})
	items := make([]engine.Item, 0)
	items = append(items, engine.Item{Id: fmt.Sprintf("%v-item", url)})
	return engine.ParseResult{Requests: requests, Items: items}
}
