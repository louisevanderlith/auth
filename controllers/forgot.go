package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Forgot struct {
}

// @Title GetForgotPage
// @Description Gets the form a user must fill in to reset their password
// @Success 200 {string} string
// @router /:forgotKey [get]
func (req *Forgot) Get(ctx context.Requester) (int, interface{}) {
	//req.Setup("forgot", "Forgot Password", true)

	return http.StatusOK, nil
}
