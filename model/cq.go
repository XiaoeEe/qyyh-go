package model

type CQMessage struct {
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	NoticeType  string `json:"notice_type"`
	Time        int64  `json:"time"`
	SelfId      int64  `json:"self_id"`
	SubType     string `json:"sub_type"`
	MessageId   int64  `json:"message_id"`
	Anonymous   bool   `json:"anonymous"`
	Font        int64  `json:"font"`
	GroupId     int64  `json:"group_id"`
	Message     string `json:"message"`
	Sender      struct {
		UserId   int64  `json:"user_id"`
		Nickname string `json:"nickname"`
		Sex      string `json:"sex"`
	} `json:"sender"`
	MessageSeq int64  `json:"message_seq"`
	RawMessage string `json:"raw_message"`
	UserId     int64  `json:"user_id"`
}

type SendGroupMessage struct {
	GroupId int64  `json:"group_id"`
	Message string `json:"message"`
}

type SendMessage struct {
	UserId  int64  `json:"user_id"`
	GroupId int64  `json:"group_id"`
	Message string `json:"message"`
}

type GroupUser struct {
	GroupId  int64  `json:"group_id"`
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	JoinTime int64  `json:"join_time"`
	Role     string `json:"role"`
}
