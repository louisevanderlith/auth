package routers

import (
	"github.com/louisevanderlith/auth/controllers"
	"github.com/louisevanderlith/droxolite/resins"
	"github.com/louisevanderlith/droxolite/roletype"
	"github.com/louisevanderlith/droxolite/routing"
)

func Setup(e resins.Epoxi) {
	homeCtrl := &controllers.Home{}
	lognCtrl := &controllers.Login{}
	regCtrl := &controllers.Register{}
	frgtCtrl := &controllers.Forgot{}
	subCtrl := &controllers.Subscribe{}

	authGroup := routing.NewInterfaceBundle("", roletype.Unknown, homeCtrl, lognCtrl, regCtrl, frgtCtrl, subCtrl)
	//q := make(map[string]string)
	//q["return"] = "{return}"
	//authGroup.AddRouteWithQueries("Login", "", "GET", roletype.Unknown, q, lognCtrl.Default)
	e.AddGroup(authGroup)
}
