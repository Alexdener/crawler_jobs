package job

import (
	"encoding/json"
	"log"
	"net/url"
	"strconv"

	"crawler_jobs/engine"
	"crawler_jobs/parser"
)

func ParseJobList(bytes []byte, fetchUrl string) (result engine.ParseResult) {
	resp := new(JobListData)
	err := json.Unmarshal(bytes, resp)
	if err != nil {
		log.Printf("ParseJobList err=%v", err)
		return
	}
	if resp.Code != 0 {
		log.Printf("ParseJobList resp=%+v", resp)
		return
	}
	for _, job := range resp.ZpData.JobList {
		result.Items = append(result.Items, engine.Item{
			Id:   job.EncryptJobID,
			Data: job,
		})
	}
	if resp.ZpData.HasMore {
		if r, err := url.Parse(fetchUrl); err == nil && r != nil {
			params := r.Query()
			if page := params.Get("page"); page != "" {
				p, _ := strconv.Atoi(page)
				params.Set("page", strconv.Itoa(p+1))
			} else {
				params.Set("page", "2")
			}
			r.RawQuery = params.Encode()
			nextFetchUrl := r.String()
			result.Requests = append(result.Requests, engine.Request{
				Url:    nextFetchUrl,
				Parser: parser.NewFuncParser(ParseJobList, "ParseJobList"),
			})
		}
	}
	return
}
