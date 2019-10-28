package entity

import (
	"polo/models"
	"polo/utils"
)

type SignInResponse struct {
	IsValid bool         `json:"isValid"`
	User    *models.User `json:"user"`
	Token   string       `json:"token"`
}

func NewSignInResponse(isValid bool, user *models.User) *SignInResponse {
	token := ""
	if isValid {
		token = utils.NewUUID()
	}
	response := SignInResponse{
		IsValid: isValid,
		User:    user,
		Token:   token,
	}
	return &response
}
