package model

type Message struct {
	MID      int64  `json:"mid"`
	OwnerUID int64  `json:"owner_uid"`
	RecUID   int64  `json:"rec_uid"`
	Detail   string `json:"detail"`
}
