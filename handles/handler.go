package handles

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/kong"
	"net/http"
)

var (
	SessionStore sessions.Store
	Authority    kong.Authority
)

func SetupRoutes(clnt, scrt, securityUrl string) http.Handler {
	stor := sessions.NewCookieStore(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32),
	)

	stor.Options.Secure = true
	stor.Options.HttpOnly = true

	SessionStore = stor

	tmpl, err := drx.LoadTemplate("./views")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.Queries("client", "{client}", "callback", "{callback}")

	r.HandleFunc("/", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, "", Index(tmpl), make(map[string]bool))).Methods(http.MethodGet)
	r.HandleFunc("/login", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, "", LoginGET(tmpl), make(map[string]bool))).Methods(http.MethodGet)
	//r.HandleFunc("/login", LoginPOST).Methods(http.MethodPost)
	r.HandleFunc("/consent", kong.ClientMiddleware(http.DefaultClient, clnt, scrt, securityUrl, "", ConsentGET(tmpl), make(map[string]bool))).Methods(http.MethodGet)
	//r.HandleFunc("/consent", ConsentPOST).Methods(http.MethodPost)

	return r
}
