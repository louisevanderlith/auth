package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Login struct {
}

func (req *Login) AcceptsQuery() map[string]string {
	q := make(map[string]string)
	q["return"] = "{return}"

	return q
}

// @Title GetLoginPage
// @Description Gets the form a user must fill in to login
// @Success 200 {string} string
// @router / [get]
func (req *Login) Get(ctx context.Requester) (int, interface{}) {
	return http.StatusOK, nil
}
