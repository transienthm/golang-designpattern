package structure

import (
	"net/http"
	"fmt"
)

//请求接口
type Request interface {
	HttpRequest() (*http.Request, error)
}

type Client struct {
	Client *http.Client
}

func (c *Client) Query(req Request) (resp *http.Response, err error) {
	httpReq, _ := req.HttpRequest()
	resp, err = c.Client.Do(httpReq)
	return
}

//实现
type CdnRequest struct {
}

func (cdn *CdnRequest) HttpRequest() (*http.Request, error) {
	return http.NewRequest("GET", "/cdn", nil)
}

//实现2
type LiveRequest struct {
}

func (cdn *LiveRequest) HttpRequest() (*http.Request, error) {
	return http.NewRequest("GET", "/live", nil)
}

func TestBridge()  {
	client := &Client{http.DefaultClient}

	cdnReq := &CdnRequest{}
	fmt.Println(client.Query(cdnReq))

	liveReq := &LiveRequest{}
	fmt.Println(client.Query(liveReq))
}