package model

import "time"

type QQParm struct {
	OpenId      string `json:"openId"`
	AccessToken string `json:"accessToken"`
}

type QQRegeditParm struct {
	QQ          int64  `json:"qq"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	AccessToken string `json:"accessToken"`
}

type LoginRet struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	QQ       int64  `json:"qq"`
}

type MapleStoryInfo struct {
	Player      []MapleStoryPlayer `json:"player"`
	SubFlagrace int                `json:"subFlagrace"`
}

type MapleStoryPlayer struct {
	Name     string    `json:"name"`
	Class    string    `json:"class"`
	Level    int64     `json:"level"`
	Img      string    `json:"img"`
	DataTime time.Time `json:"dataTime"`
	Points   int       `json:"points"`
	Culvert  int       `json:"culvert"`
	Flagrace int       `json:"flagrace"`
}
