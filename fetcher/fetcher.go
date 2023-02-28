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
