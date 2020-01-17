package SwagDoc

import (
	"GoSwaggerConvert/Req"
	"GoSwaggerConvert/Resp"
)

// 好友資料 - 回傳資料
//
// swagger:response  FriendGetResp
type FriendGetResp struct {
	// in:body
	Body struct {
		Resp.BaseResp
		Data struct {
			Resp.FriendGetResp
		}
	} `json:"resp_data"`
}

// 好友請求資料
//
// swagger:parameters  FriendGet
type FriendGetReq struct {
	// in:body
	Body struct {
		Body Req.FriendGetReq `json:"req_data"`
	}
}
