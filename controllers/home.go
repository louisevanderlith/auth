package controllers

import (
	"net/http"

	"github.com/louisevanderlith/droxolite/context"
)

type Home struct {
}

func (r *Home) Default(ctx context.Contexer) (int, interface{}) {
	return http.StatusOK, nil
}
