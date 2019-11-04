package client

import (
	"log"
	"net/http"
)

//New Clientを生成します
func New(httpclient http.Client) Client {
	return &client{httpclient}
}

//Client 各種特殊なクライアントを提供します
type Client interface {
}

//Client 内部の情報を定義します
type client struct {
	httpclient http.Client
}

func (c *client) Do(req *http.Request) (resp *http.Response, err error) {
	resp, err = c.httpclient.Do(req)
	if err != nil {
		log.Panicln(err)
		return resp, err
	}
	defer resp.Body.Close()
	return resp, err
}
