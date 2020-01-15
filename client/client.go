package client

import (
	"log"
	"net/http"
	"time"

	"github.com/MatchlockIsDAST/httpmesconv/stringto"
	"github.com/MatchlockIsDAST/httpmesconv/tostring"

	"github.com/MatchlockIsDAST/sphttpclient/judgment"
)

//New Clientを生成します
func New(httpclient http.Client) Client {
	return &client{httpclient}
}

//Client 各種特殊なクライアントを提供します
type Client interface {
	ChangeClient(httpclient http.Client)
	GetClient() http.Client
	Do(req *http.Request) (resp *http.Response, err error)
	TimeBaseJudgeDo(elapsedMin, elapsedMax time.Duration, req *http.Request) (resp *http.Response, flag bool, err error)
	DisplayBaseJudgeDo(included string, req *http.Request) (resp *http.Response, flag bool, err error)
	DiffBaseJudgeDo(shouldbe bool, requests []*http.Request) (resps []*http.Response, flag bool, err error)
}

//Client 内部の情報を定義します
type client struct {
	httpclient http.Client
}

func (c *client) ChangeClient(httpclient http.Client) {
	c.httpclient = httpclient
}

func (c *client) GetClient() http.Client {
	return c.httpclient
}

//通常のClientを提供する
func (c *client) Do(req *http.Request) (resp *http.Response, err error) {
	resp, err = c.httpclient.Do(req)
	if err != nil {
		log.Panicln(err)
		return resp, err
	}
	return resp, err
}

//時間計測を同時に行うClientを提供する
func (c *client) TimeBaseJudgeDo(elapsedMin, elapsedMax time.Duration, req *http.Request) (resp *http.Response, flag bool, err error) {
	start := time.Now()
	resp, err = c.Do(req)
	duration := time.Now().Sub(start)
	if err != nil {
		return nil, false, err
	}
	flag = judgment.TimeBase(elapsedMin, elapsedMax, duration)
	return resp, flag, err
}

//表示判定を行うClientを提供する
func (c *client) DisplayBaseJudgeDo(included string, req *http.Request) (resp *http.Response, flag bool, err error) {
	resp, err = c.Do(req)
	if err != nil {
		return nil, false, err
	}
	body := tostring.Body(resp.Body)
	flag = judgment.DisplayBase(body, included)
	resp.Body = stringto.IoReadCloser(body)
	return resp, flag, err
}

//差分判定を行うclient
//shouldbe		: 完全一致が正常 True , 完全不一致が正常 False
//requests		: 検証用のHTTPリクエスト
func (c *client) DiffBaseJudgeDo(shouldbe bool, requests []*http.Request) (resps []*http.Response, flag bool, err error) {
	resps = make([]*http.Response, len(requests))
	bodys := make([]string, len(requests))
	for i := 0; i < len(requests); i++ {
		resps[i], err = c.Do(requests[i])
		bodys[i] = tostring.Body(resps[i].Body)
		resps[i].Body = stringto.IoReadCloser(bodys[i])
		if err != nil {
			return nil, false, err
		}
	}
	flag = judgment.DiffBase(bodys)
	//flag == shouldbe 差分一致の判定と差分可否をandすると結果が確認できる
	return resps, flag == shouldbe, nil
}
