package main

import (
	"flag"
	"github.com/louisevanderlith/auth/handles"
	"github.com/louisevanderlith/kong"
	"net/http"
	"time"

	"github.com/louisevanderlith/droxolite"
)

func main() {
	clientId := flag.String("client", "mango.auth", "Client ID which will be used to verify this instance")
	clientSecrt := flag.String("secret", "secret", "Client Secret which will be used to authenticate this instance")
	authrty := flag.String("authority", "http://localhost:8086", "Authority Provider's URL")

	flag.Parse()

	tkn, err := kong.FetchToken(http.DefaultClient, *authrty, *clientId, *clientSecrt, "theme.assets.download", "theme.assets.view")

	if err != nil {
		panic(err)
	}

	clms, err := kong.Exchange(tkn, *clientId, *clientSecrt, *authrty+"/info")

	if err != nil {
		panic(err)
	}

	err = droxolite.UpdateTemplate(clms)

	if err != nil {
		panic(err)
	}

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8084",
		Handler:      handles.SetupRoutes(),
	}

	err = srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
