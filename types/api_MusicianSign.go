package types

// MusicianSignData 音乐人签到返回数据
type MusicianSignData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    bool   `json:"data"`
}
