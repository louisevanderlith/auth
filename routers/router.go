// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/louisevanderlith/auth/controllers"
	"github.com/louisevanderlith/mango"

	"github.com/astaxie/beego"
	"github.com/louisevanderlith/mango/control"
	secure "github.com/louisevanderlith/secure/core"
)

func Setup(s *mango.Service) {
	ctrlmap := EnableFilter(s)

	siteName := beego.AppConfig.String("defaultsite")
	theme, err := mango.GetDefaultTheme(ctrlmap.GetInstanceID(), siteName)

	if err != nil {
		panic(err)
	}

	lognCtrl := controllers.NewLoginCtrl(ctrlmap, theme)

	beego.Router("/login", lognCtrl, "get:Get")
	beego.Router("/register", controllers.NewRegisterCtrl(ctrlmap, theme), "get:Get")
	beego.Router("/subscribe", controllers.NewSubscribeCtrl(ctrlmap, theme), "get:Get")

	beego.Router("/forgot", controllers.NewForgotCtrl(ctrlmap, theme), "get:Get")
}

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
