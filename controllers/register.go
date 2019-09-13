package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Register struct {
}

// @Title GetRegisterPage
// @Description Gets the form a user must fill in to register
// @Success 200 {string} string
// @router / [get]
func (req *Register) Get(ctx context.Requester) (int, interface{}) {
	return http.StatusOK, nil
}
