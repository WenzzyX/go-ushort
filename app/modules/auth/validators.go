package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/wenzzyx/go-ushort/app/common/utils"
	"github.com/wenzzyx/go-ushort/app/models"
)

type RegisterValidator struct {
	Email     string           `json:"email" binding:"required,email"`
	Password  string           `json:"password" binding:"required,min=8,max=64"`
	UserModel models.UserModel `json:"-"`
}

func (v *RegisterValidator) Bind(c *gin.Context) error {
	if err := utils.Bind(c, v); err != nil {
		return err
	}
	v.UserModel.Email = v.Email
	if err := v.UserModel.SetPassword(v.Password); err != nil {
		return err
	}
	return nil
}

func NewRegisterValidator() RegisterValidator {
	v := RegisterValidator{}
	return v
}

// ###

type LoginValidator struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (v *LoginValidator) Bind(c *gin.Context) error {
	if err := utils.Bind(c, v); err != nil {
		return err
	}
	return nil
}

func NewLoginValidator() LoginValidator {
	v := LoginValidator{}
	return v
}

// ###

type RefreshValidator struct {
	Token string `json:"token" binding:""`
}

func (v *RefreshValidator) Bind(c *gin.Context) error {
	if err := utils.Bind(c, v); err != nil {
		return err
	}
	return nil
}

func NewRefreshValidator() RefreshValidator {
	v := RefreshValidator{}
	return v
}
