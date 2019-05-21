package controllers

import (
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/control"
)

type SubscribeController struct {
	control.UIController
}

func NewSubscribeCtrl(ctrlMap *control.ControllerMap, setting mango.ThemeSetting) *SubscribeController {
	result := &SubscribeController{}
	result.SetTheme(setting)
	result.SetInstanceMap(ctrlMap)

	return result
}

// @Title GetSubribePage
// @Description Gets the form a user must fill in to allow a service
// @Success 200 {string} string
// @router / [get]
func (req *SubscribeController) Get() {
	req.Setup("subscribe", "Subscribe", false)
}
