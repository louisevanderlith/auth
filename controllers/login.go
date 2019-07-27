package controllers

import "github.com/louisevanderlith/droxolite/xontrols"

type LoginController struct {
	xontrols.UICtrl
}

// @Title GetLoginPage
// @Description Gets the form a user must fill in to login
// @Success 200 {string} string
// @router / [get]
func (req *LoginController) Get() {
	req.Setup("login", "Login", true)
}
