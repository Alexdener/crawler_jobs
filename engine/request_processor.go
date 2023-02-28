package engine

import "crawler_jobs/fetcher"

func RequestProcess(r Request) (result ParseResult, err error) {
	content, err := fetcher.Fetch(r.Url)
	if err != nil {
		return
	}

	return r.Parser.Parse(content, r.Url), nil
}
