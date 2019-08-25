package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Subscribe struct {
}

// @Title GetSubribePage
// @Description Gets the form a user must fill in to allow a service
// @Success 200 {string} string
// @router / [get]
func (req *Subscribe) Default(ctx context.Contexer) (int, interface{}) {
	//req.Setup("subscribe", "Subscribe", false)

	return http.StatusOK, nil
}
