package nldyp

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
)

type PartnerData4SellTicketResponse struct {
	Status int `json:"status"`
	Data   []struct {
		OrderId string `json:"orderId"` // 订单 id
		OrderNo string `json:"orderNo"` // 系统商锁座订单号
	} `json:"data"`
	Content string `json:"content"`
}

type PartnerData4SellTicketResult struct {
	Result PartnerData4SellTicketResponse // 结果
	Body   []byte                         // 内容
	Http   gorequest.Response             // 请求
}

func newPartnerData4SellTicketResult(result PartnerData4SellTicketResponse, body []byte, http gorequest.Response) *PartnerData4SellTicketResult {
	return &PartnerData4SellTicketResult{Result: result, Body: body, Http: http}
}

// PartnerData4SellTicket 售票
// https://docs.apipost.cn/preview/fa101f4865dc783f/66e7c2e894fda4a6?target_id=edd25537-3425-49e0-b6b5-373ebe4e919c
func (c *Client) PartnerData4SellTicket(ctx context.Context, orderId, orderNo, notifyUrl string, notMustParams ...gorequest.Params) (*PartnerData4SellTicketResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("orderId", orderId) // 订单 id
	if orderNo != "" {
		params.Set("orderNo", orderNo) // 接入方平台订单号（方便快速查询平台订单信息，又接入方自行生成）
	}
	params.Set("notifyUrl", notifyUrl) // 退票通知地址
	// 请求
	request, err := c.request(ctx, apiUrl+"/partner/data4/sellTicket", params)
	if err != nil {
		return newPartnerData4SellTicketResult(PartnerData4SellTicketResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PartnerData4SellTicketResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPartnerData4SellTicketResult(response, request.ResponseBody, request), err
}
