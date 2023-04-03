package fetcher

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) (content []byte, err error) {
	resp, err := http.Get(url)
	log.Println("url", url, "resp", resp.StatusCode, "err", err)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		return
	}
	return ioutil.ReadAll(resp.Body)
}

func JobFetch(url string) (content []byte, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header.Add("user-agent", "TODO") // TODO
	req.Header.Add("cookie", "TODO;")    // TODO
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer func() {
		log.Println("url", url, "header", req.Header, "resp", resp.StatusCode, "content", string(content), "err", err)
	}()
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = errors.New(resp.Status)
		return
	}
	content, err = ioutil.ReadAll(resp.Body)
	return
}
