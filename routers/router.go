package routers

import (
	"github.com/louisevanderlith/auth/controllers"
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/droxolite/roletype"
)

func Setup(poxy *droxolite.Epoxy) {
	//Forgot
	frgtCtrl := &controllers.ForgotController{}
	frgtGroup := droxolite.NewRouteGroup("forgot", frgtCtrl)
	frgtGroup.AddRoute("/", "GET", roletype.Unknown, frgtCtrl.Get)
	poxy.AddGroup(frgtGroup)

	//Login
	lognCtrl := &controllers.LoginController{}
	lognGroup := droxolite.NewRouteGroup("login", lognCtrl)
	lognGroup.AddRoute("/", "GET", roletype.Unknown, lognCtrl.Get)
	poxy.AddGroup(lognGroup)

	//Register
	regCtrl := &controllers.RegisterController{}
	regGroup := droxolite.NewRouteGroup("register", regCtrl)
	regGroup.AddRoute("/", "GET", roletype.Unknown, regCtrl.Get)
	poxy.AddGroup(regGroup)

	//Subscribe
	subCtrl := &controllers.SubscribeController{}
	subGroup := droxolite.NewRouteGroup("upload", subCtrl)
	subGroup.AddRoute("/", "GET", roletype.Unknown, subCtrl.Get)
	poxy.AddGroup(frgtGroup)

	/*ctrlmap := EnableFilter(s)

	siteName := beego.AppConfig.String("defaultsite")
	theme, err := mango.GetDefaultTheme(ctrlmap.GetInstanceID(), siteName)

	if err != nil {
		panic(err)
	}

	lognCtrl := controllers.NewLoginCtrl(ctrlmap, theme)

	beego.Router("/login", lognCtrl, "get:Get")
	beego.Router("/register", controllers.NewRegisterCtrl(ctrlmap, theme), "get:Get")
	beego.Router("/subscribe", controllers.NewSubscribeCtrl(ctrlmap, theme), "get:Get")

	beego.Router("/forgot", controllers.NewForgotCtrl(ctrlmap, theme), "get:Get")*/
}

/*
func EnableFilter(s *mango.Service) *control.ControllerMap {
	ctrlmap := control.CreateControlMap(s)

	emptyMap := make(secure.ActionMap)
	ctrlmap.Add("/subscribe", emptyMap)
	ctrlmap.Add("/login", emptyMap)
	ctrlmap.Add("/register", emptyMap)
	ctrlmap.Add("/forgot", emptyMap)

	beego.InsertFilter("/*", beego.BeforeRouter, ctrlmap.FilterUI)

	return ctrlmap
}
*/
