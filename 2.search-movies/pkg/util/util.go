package util

import (
	"io/ioutil"
	"net/http"
)

// DoHTTPRequest do a http request
func DoHTTPRequest(outgoingReq *http.Request) (body []byte, status int, err error) {
	client := &http.Client{}
	res, err := client.Do(outgoingReq)
	if err != nil {
		return
	}
	defer res.Body.Close()

	status = res.StatusCode
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	return
}
