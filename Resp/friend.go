package Resp

type FriendGetResp struct {
	// uid
	// example:0004a15b
	User_id string `json:"user_id"`
	// 名稱
	// example:apple
	User_name string `json:"user_name"`
	// alias
	// example:llllkkkkk
	User_alias string `json:"user_alias"`
	// 照片
	// example:default
	User_image string `json:"user_image"`
	// 禁言
	// example:0
	Is_mute string `json:"is_mute"`
	// 最上層
	// example:0
	Is_top string `json:"is_top"`
	// 狀態
	// example:0
	Status string `json:"status"`
}
