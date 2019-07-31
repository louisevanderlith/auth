package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/xontrols"
)

type ForgotController struct {
	xontrols.UICtrl
}

// @Title GetForgotPage
// @Description Gets the form a user must fill in to reset their password
// @Success 200 {string} string
// @router /:forgotKey [get]
func (req *ForgotController) Get() {
	req.Setup("forgot", "Forgot Password", true)

	req.Serve(http.StatusOK, nil, nil)
}
