package widget

type Token struct {
	Token    string `json:"token"`     //小部件token
	ExpireAt int64  `json:"expire_at"` //过期时间
}
