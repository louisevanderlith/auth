package controllers

import "github.com/louisevanderlith/mango/control"

type LoginController struct {
	control.UIController
}

func NewLoginCtrl(ctrlMap *control.ControllerMap) *LoginController {
	result := &LoginController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetLoginPage
// @Description Gets the form a user must fill in to login
// @Success 200 {string} string
// @router / [get]
func (req *LoginController) Get() {
	req.Setup("login")
}
