package Req

type FriendGetReq struct {
	// 時間戳
	// example:1567765485000
	Time string `json:"time"`
	// 用戶ID
	// exampie:e5d8r33t9
	UserID string `json:"user_id"`
}
