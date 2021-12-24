package util

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type NetworkManager struct {
}

type NetworkRequest struct {
	owner    string
	req      *http.Request
	timeout  int
	complete context.CancelFunc
	method   string
	url      string
}

func (self *NetworkRequest) AddHeader(headerName string, headerValue string) {
	self.req.Header.Add(headerName, headerValue)
}

func (self *NetworkRequest) AddQueryParam(key string, value string) {

	//
	q := self.req.URL.Query()
	q.Add(key, value)

	//
	self.req.URL.RawQuery = q.Encode()

	//
	log.Println("AddQueryParam: URL = %s", self.req.URL.String())

}

func NewNetworkManager() *NetworkManager {
	return new(NetworkManager)
}

func (self *NetworkManager) MakeRequest(owner string, method string, url string, timeout int) (*NetworkRequest, error) {
	newNetworkRequest := new(NetworkRequest)
	newNetworkRequest.owner = owner
	newNetworkRequest.timeout = timeout
	newNetworkRequest.method = method
	newNetworkRequest.url = url
	contextDuration, _ := time.ParseDuration(fmt.Sprintf("%ds", timeout))
	parentContext := context.Background()
	c, complete := context.WithTimeout(parentContext, contextDuration)
	newNetworkRequest.complete = complete
	req, err := http.NewRequestWithContext(c, method, url, nil)
	newNetworkRequest.req = req
	return newNetworkRequest, err
}

type NetworkResponse struct {
	Body []byte
}

func (self *NetworkResponse) Close() {

}

func (self *NetworkResponse) Bytes() []byte {
	return self.Body
}

func (self *NetworkManager) Perform(networkRequest *NetworkRequest) (*NetworkResponse, error) {
	log.Printf("Make HTTP request: owner = %s method = %s url = %s timeout = %d sec. ",
		networkRequest.owner,
		networkRequest.method,
		networkRequest.url,
		networkRequest.timeout,
	)

	/* Step 1. Download response */
	client := http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Do(networkRequest.req)
	if err != nil {
		log.Printf("HTTP error: err = %#v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var out bytes.Buffer
	io.Copy(&out, resp.Body)
	log.Printf("out = %s", out.String())

	/* Complete */
	networkRequest.complete()

	newNetworkResponse := new(NetworkResponse)
	newNetworkResponse.Body = out.Bytes()

	return newNetworkResponse, nil
}
