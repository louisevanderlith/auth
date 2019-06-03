package main

import (
	"log"
	"os"
	"path"

	"github.com/louisevanderlith/auth/routers"
	"github.com/louisevanderlith/mango"
	"github.com/louisevanderlith/mango/enums"

	"github.com/astaxie/beego"
)

func main() {
	keyPath := os.Getenv("KEYPATH")
	pubName := os.Getenv("PUBLICKEY")
	pubPath := path.Join(keyPath, pubName)

	// Register with router
	appName := beego.BConfig.AppName
	srv := mango.NewService(appName, pubPath, enums.APP)

	port := beego.AppConfig.String("httpport")
	err := srv.Register(port)

	if err != nil {
		log.Print("Register: ", err)
	} else {
		err = mango.UpdateTheme(srv.ID)

		if err != nil {
			panic(err)
		}

		routers.Setup(srv)

		beego.SetStaticPath("/dist", "dist")
		beego.Run()
	}
}
