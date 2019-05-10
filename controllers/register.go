package controllers

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type RegisterController struct {
	control.UIController
}

func NewRegisterCtrl(ctrlMap *control.ControllerMap, setting mango.ThemeSetting) *RegisterController {
	result := &RegisterController{}
	result.SetTheme(setting)
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetRegisterPage
// @Description Gets the form a user must fill in to register
// @Success 200 {string} string
// @router / [get]
func (req *RegisterController) Get() {
	req.Setup("register", "Registration", true)
}
