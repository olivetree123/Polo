package models

import (
	"polo/utils"
)

type User struct {
	BaseModel
	Account string
	Cipher  string
	Avatar  string
}

// NewUser 新建用户
func NewUser(account string, cipher string) (*User, error) {
	user := User{
		Account: account,
		Cipher:  utils.ContentMD5([]byte(cipher)),
		Avatar:  "",
	}
	err := DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ValidateCipher 验证密码
func ValidateCipher(account string, cipher string) (bool, *User, error) {
	var user User
	cipher2 := utils.ContentMD5([]byte(cipher))
	err := DB.First(&user, "account = ?", account).Error
	if err != nil {
		return false, nil, err
	}
	if user.Cipher == cipher2 {
		return true, &user, nil
	}
	return false, nil, nil
}
