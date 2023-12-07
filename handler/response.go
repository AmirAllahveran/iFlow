package handler

import (
	"github.com/smf8/http-monitor/common"
	"github.com/smf8/http-monitor/model"
)

type ResponseData struct {
	Data interface{} `json:"data"`
}

func NewResponseData(data interface{}) *ResponseData {
	return &ResponseData{data}
}

type UserResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func NewUserResponse(user *model.User) *UserResponse {
	token, _ := common.GenerateJWT(user.ID)
	ur := &UserResponse{Username: user.Username, Token: token}
	return ur
}
