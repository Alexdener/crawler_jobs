package engine

import (
	"time"

	"crawler_jobs/fetcher"
)

var rateLimiter = time.Tick(5 * time.Second)

func RequestProcess(r Request) (result ParseResult, err error) {
	content, err := fetcher.Fetch(r.Url)
	if err != nil {
		return
	}

	return r.Parser.Parse(content, r.Url), nil
}

func JobListRequestProcess(r Request) (result ParseResult, err error) {
	<-rateLimiter
	content, err := fetcher.JobFetch(r.Url)
	if err != nil {
		return
	}

	return r.Parser.Parse(content, r.Url), nil
}
