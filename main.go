package main

import (
	"crawler_jobs/engine"
	"crawler_jobs/parser"
	"crawler_jobs/parser/mock"
	"crawler_jobs/scheduler"
)

func main() {
	e := engine.Engine{
		Scheduler:        &scheduler.QueueScheduler{},
		WorkerCount:      3,
		RequestProcessor: engine.RequestProcess,
		ItemChan:         make(chan engine.Item),
	}
	e.Run(engine.Request{
		Url:    "http://127.0.0.1:40001/mock/11/webA",
		Parser: parser.NewFuncParser(mock.ParseMockRequest, "ParseMockRequest"),
	})
}
