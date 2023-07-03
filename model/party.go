package model

type GetPartyListParm struct {
	Type int `json:"type"`
}

type CreatePartyParm struct {
	Boss    []string `json:"boss"`
	Leader  string   `json:"leader"`
	Count   int      `json:"count"`
	Level   int64    `json:"level"`
	Point   int      `json:"point"`
	Remarks string   `json:"remarks"`
	Time    string   `json:"time"`
	Type    int      `json:"type"`
	Channel int      `json:"channel"`
}

type GetPartyListReq struct {
	ID       int64         `json:"id"`
	Title    string        `json:"title"`
	Type     int           `json:"type"`
	Remarks  string        `json:"remarks" `
	Level    int64         `json:"level"`
	Time     string        `json:"time"`
	Ch       int           `json:"ch"`
	Point    int           `json:"point"`
	Leader   PartyPlayer   `json:"leader"`
	Member   []PartyPlayer `json:"member"`
	CreateBy int           `json:"createBy"`
}

type PartyPlayer struct {
	QQ       int64  `json:"qq"`
	Name     string `json:"name"`
	Level    int64  `json:"level"`
	Job      string `json:"job"`
	Point    int    `json:"point"`
	UsePoint int    `json:"usePoint"`
	Img      string `json:"img"`
}

type DelPartyParm struct {
	ID int64 `json:"id"`
}

type JoinPartyParm struct {
	ID int64  `json:"id"`
	Me string `json:"me"`
}

type LeavePartyParm struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
