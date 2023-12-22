package nldyp

import (
	"context"
	"go.dtapp.net/gorequest"
)

func (c *Client) request(ctx context.Context, url string, param gorequest.Params) (gorequest.Response, error) {

	// 签名
	params := c.Sign(param)

	// 创建请求
	client := gorequest.NewHttp()
	if !c.requestClientStatus {
		c.DefaultHttp()
		client = c.requestClient
	}

	// 设置请求地址
	client.SetUri(url)

	// 设置格式
	client.SetContentTypeJson()

	// 设置用户代理
	client.SetUserAgent(gorequest.GetRandomUserAgentSystem())

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Post(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.gormLog.status {
		go c.gormLog.client.Middleware(ctx, request)
	}

	return request, err
}
