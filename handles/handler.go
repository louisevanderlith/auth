package handles

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/louisevanderlith/droxolite"
	"github.com/louisevanderlith/kong"
	"net/http"
)

var (
	SessionStore sessions.Store
	Security     kong.Securer
)

func SetupRoutes(clnt, scrt, secureUrl string) http.Handler {
	stor := sessions.NewCookieStore(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32),
	)

	stor.Options.Secure = true
	stor.Options.HttpOnly = true

	SessionStore = stor

	mstr, tmpl, err := droxolite.LoadTemplate("./views", "master.html")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.Queries("client", "{client}", "callback", "{callback}")

	r.HandleFunc("/login", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, "", LoginGET(mstr, tmpl))).Methods(http.MethodGet)
	r.HandleFunc("/login", LoginPOST).Methods(http.MethodPost)
	r.HandleFunc("/consent", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, secureUrl, "", ConsentGET(mstr, tmpl))).Methods(http.MethodGet)
	r.HandleFunc("/consent", ConsentPOST).Methods(http.MethodPost)

	return r
}
