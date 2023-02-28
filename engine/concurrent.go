package engine

import "log"

type Engine struct {
	Scheduler        Scheduler //调度器
	WorkerCount      int       //worker数量
	RequestProcessor Processor //请求处理器，分布式爬虫可实现rpc调用
	ItemChan         chan Item //数据存储
}

func (e *Engine) Run(seeds ...Request) {
	//等待request和worker 1个request分配给1个worker
	e.Scheduler.Run()
	for _, seed := range seeds {
		e.Scheduler.Submit(seed)
	}

	out := make(chan ParseResult)
	for i := 0; i < e.WorkerCount; i++ {
		go e.createWorker(e.Scheduler.CreateWorkerChan(), out)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Println("save item", item.Id)
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func (e *Engine) createWorker(in chan Request, out chan ParseResult) {
	for {
		e.Scheduler.WorkerReady(in)
		r := <-in
		parseResult, err := e.RequestProcessor(r)
		if err != nil {
			log.Println("RequestProcessor error", err)
			continue
		}
		out <- parseResult
		//select {
		//case r := <- in:
		//	parseResult, err := e.RequestProcessor(r)
		//	if err != nil {
		//		log.Println("RequestProcessor error", err)
		//		continue
		//	}
		//	out <- parseResult
		//}
	}
}
