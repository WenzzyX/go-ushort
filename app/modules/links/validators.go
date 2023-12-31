package links

import (
	"github.com/gin-gonic/gin"

	"github.com/wenzzyx/go-ushort/app/common/utils"
)

type LinkCreateValidator struct {
	Name    string `json:"name" binding:"max=150"`
	RealUrl string `json:"realUrl" binding:"required,url,min=5,max=2000"`
}

func (v *LinkCreateValidator) Bind(c *gin.Context) error {
	if err := utils.Bind(c, v); err != nil {
		return err
	}
	return nil
}

func NewLinkCreateValidator() LinkCreateValidator {
	v := LinkCreateValidator{}
	return v
}

type RedirectUriValidator struct {
	Alias string `json:"alias" uri:"alias"`
}

type LinkUpdateUriValidator struct {
	ID uint `json:"id" uri:"linkId"`
}
type LinkUpdateValidator struct {
	Name string `json:"name" binding:"max=150"`
}

func (v *LinkUpdateValidator) Bind(c *gin.Context) error {
	if err := utils.Bind(c, v); err != nil {
		return err
	}
	return nil
}

func NewLinkUpdateValidator() LinkUpdateValidator {
	v := LinkUpdateValidator{}
	return v
}
