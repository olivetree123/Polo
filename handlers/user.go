package handlers

import (
	"encoding/json"
	"github.com/olivetree123/coco"
	"polo/entity"
	"polo/models"
	"polo/utils"
)

type SignUpParam struct {
	Account string
	Cipher  string
}

type SignInParam SignUpParam

type SignInResponse struct {
	IsValid bool         `json:"isValid"`
	User    *models.User `json:"user"`
	Token   string       `json:"token"`
}

// SignUpHandler 注册
func SignUpHandler(c *coco.Coco) coco.Result {
	var param SignUpParam
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&param)
	if err != nil {
		panic(err)
	}
	//c.GetJsonParam(&)
	user, err := models.NewUser(param.Account, param.Cipher)
	if err != nil {
		panic(err)
	}
	return coco.APIResponse(user)
}

// SignInHandler 登陆
func SignInHandler(c *coco.Coco) coco.Result {
	var param SignInParam
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&param)
	if err != nil {
		panic(err)
	}
	isValid, user, err := models.ValidateCipher(param.Account, param.Cipher)
	if err != nil {
		panic(err)
	}
	result := entity.NewSignInResponse(isValid, user)
	err = utils.SetCache(result.Token, "1", 3600)
	if err != nil {
		panic(err)
	}
	return coco.APIResponse(result)
}

// SignOutHandler 登出
//func SignOutHandler(c *coco.Coco) coco.Result {
//	token := c.Request.Header.Get("Authorization")
//	jwt.Parse(token)
//}
