package controllers

import "github.com/louisevanderlith/droxolite/xontrols"

type RegisterController struct {
	xontrols.UICtrl
}

// @Title GetRegisterPage
// @Description Gets the form a user must fill in to register
// @Success 200 {string} string
// @router / [get]
func (req *RegisterController) Get() {
	req.Setup("register", "Registration", true)
}
