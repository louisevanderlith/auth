package main

import (
	"flag"
	"github.com/louisevanderlith/auth/handles"
	"net/http"
	"time"

	"github.com/louisevanderlith/droxolite"
)

func main() {
	clientId := flag.String("client", "mango.auth", "Client ID which will be used to verify this instance")
	clientSecrt := flag.String("secret", "secret", "Client Secret which will be used to authenticate this instance")
	security := flag.String("security", "http://localhost:8086", "Security Provider's URL")

	flag.Parse()

	err := droxolite.UpdateTemplate(*clientId, *clientSecrt, *security)

	if err != nil {
		panic(err)
	}

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8094",
		Handler:      handles.SetupRoutes(*clientId, *clientSecrt, *security),
	}

	err = srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
