package nldyp

import (
	"context"
	"go.dtapp.net/gojson"
	"go.dtapp.net/gorequest"
)

type PartnerData4GetPlanResponse struct {
	Code int `json:"code"`
	Data []struct {
		FeatureAppNo  string `json:"featureAppNo"` // 排期编码
		CinemaCode    int    `json:"cinemaCode"`   // 影城编码
		SourceFilmNo  string `json:"sourceFilmNo"`
		FilmNo        string `json:"filmNo"`        // 影片编码
		FilmName      string `json:"filmName"`      // 影片名称
		HallNo        string `json:"hallNo"`        // 影厅编码
		HallName      string `json:"hallName"`      // 影厅名称
		StartTime     int    `json:"startTime"`     // 放映时间
		CopyType      string `json:"copyType"`      // 影片制式
		CopyLanguage  string `json:"copyLanguage"`  // 语言
		TotalTime     string `json:"totalTime"`     // 时长
		ListingPrice  string `json:"listingPrice"`  // 挂牌价
		TicketPrice   string `json:"ticketPrice"`   // 票价
		ServiceAddFee string `json:"serviceAddFee"` // 服务费下限
		LowestPrice   string `json:"lowestPrice"`   // 最低保护价
		Thresholds    string `json:"thresholds"`    // 服务费上限
		Areas         string `json:"areas"`         // 座区属性
		MarketPrice   string `json:"marketPrice"`   // 市场参考价（无座位区间时，下特价票使用该价格，有座位区间则使用座位区间价）
	} `json:"data"`
	Content string `json:"content"`
}

type PartnerData4GetPlanResult struct {
	Result PartnerData4GetPlanResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
}

func newPartnerData4GetPlanResult(result PartnerData4GetPlanResponse, body []byte, http gorequest.Response) *PartnerData4GetPlanResult {
	return &PartnerData4GetPlanResult{Result: result, Body: body, Http: http}
}

// PartnerData4GetPlan 获取影院排期
// https://docs.apipost.cn/preview/fa101f4865dc783f/66e7c2e894fda4a6?target_id=1e0161fe-12ca-4a64-afb9-5c97431f1f70
func (c *Client) PartnerData4GetPlan(ctx context.Context, notMustParams ...gorequest.Params) (*PartnerData4GetPlanResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/partner/data4/getPlan", params)
	if err != nil {
		return newPartnerData4GetPlanResult(PartnerData4GetPlanResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response PartnerData4GetPlanResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPartnerData4GetPlanResult(response, request.ResponseBody, request), err
}
