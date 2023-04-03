package main

import (
	"log"

	"crawler_jobs/engine"
	"crawler_jobs/initialize"
	"crawler_jobs/parser"
	"crawler_jobs/parser/job"
	"crawler_jobs/persist"
	"crawler_jobs/scheduler"
)

func main() {
	initialize.InitDb()
	itemChan, err := persist.SaveJobList()
	if err != nil {
		log.Fatal(err)
	}
	e := engine.Engine{
		Scheduler:        &scheduler.QueueScheduler{},
		WorkerCount:      3,
		RequestProcessor: engine.JobListRequestProcess,
		ItemChan:         itemChan,
	}
	e.Run(engine.Request{
		Url:    "http://127.0.0.1:40001/mock/11/webA?page=1",
		Parser: parser.NewFuncParser(job.ParseJobList, "ParseJobList"),
	})
}
