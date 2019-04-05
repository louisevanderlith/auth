package controllers

import (
	"github.com/louisevanderlith/mango/control"
)

type ForgotController struct {
	control.UIController
}

func NewForgotCtrl(ctrlMap *control.ControllerMap) *ForgotController {
	result := &ForgotController{}
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetForgotPage
// @Description Gets the form a user must fill in to reset their password
// @Success 200 {string} string
// @router / [get]
func (req *ForgotController) Get() {
	req.Setup("forgot", "Forgot Password", false)
}
