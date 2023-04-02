package engine

type Processor func(Request) (ParseResult, error)

type Request struct {
	Url    string
	Parser Parser
}

type Item struct {
	Id   string
	Data interface{}
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Scheduler interface {
	Submit(Request)
	CreateWorkerChan() chan Request
	WorkerReady(chan Request)
	Run()
}

type ParserFunc func([]byte, string) ParseResult

type Parser interface {
	Parse([]byte, string) ParseResult
}
