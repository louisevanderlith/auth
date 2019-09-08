package routers

import (
	"github.com/louisevanderlith/auth/controllers"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(e resins.Epoxi) {
	lognCtrl := &controllers.Login{}
	regCtrl := &controllers.Register{}
	frgtCtrl := &controllers.Forgot{}
	subCtrl := &controllers.Subscribe{}

	e.JoinBundle("/", roletype.Unknown, mix.Page, lognCtrl, regCtrl, frgtCtrl, subCtrl)
}
