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
func (req *Subscribe) Get(ctx context.Requester) (int, interface{}) {
	return http.StatusOK, nil
}
