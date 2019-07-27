package controllers

import "github.com/louisevanderlith/droxolite/xontrols"

type SubscribeController struct {
	xontrols.UICtrl
}

// @Title GetSubribePage
// @Description Gets the form a user must fill in to allow a service
// @Success 200 {string} string
// @router / [get]
func (req *SubscribeController) Get() {
	req.Setup("subscribe", "Subscribe", false)
}
