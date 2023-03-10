// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

type RegisterReq struct {
	Name     string `json:"name"`
	Gender   int64  `json:"gender"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type RegisterResp struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Gender int64  `json:"gender"`
	Mobile string `json:"mobile"`
}

type UserInfoResp struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Gender int64  `json:"gender"`
	Mobile string `json:"mobile"`
}
