package model

type BilibiliData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		LiveRoom struct {
			RoomStatus int    `json:"roomStatus"` //直播间状态 1:正常
			LiveStatus int    `json:"liveStatus"` //直播状态 0:未开播 1:正在直播
			Url        string `json:"url"`
			Title      string `json:"title"`
		} `json:"live_room"`
	} `json:"data"`
}
