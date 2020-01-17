package Resp

type BaseResp struct {
	// 狀態
	// example:Success
	StatusValue string `json:"s"`
	// 訊息
	// exmaple:請求成功
	Msg string `json:"m"`
}
