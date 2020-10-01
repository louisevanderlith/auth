package handles

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/menu"
	"github.com/louisevanderlith/kong"
	"github.com/louisevanderlith/kong/middle"
	"net/http"
)

var (
	SessionStore sessions.Store
	Authority    kong.Authority
)

func FullMenu() *menu.Menu {
	m := menu.NewMenu()

	m.AddItem(menu.NewItem("a", "/", "Home", nil))
	m.AddItem(menu.NewItem("b", "/register", "Register", nil))
	m.AddItem(menu.NewItem("c", "/forgot", "Forgot Password", nil))

	return m
}

func SetupRoutes(clnt, scrt, securityUrl, managerUrl string) http.Handler {
	stor := sessions.NewCookieStore(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32),
	)

	//stor.Options.Secure = true
	stor.Options.HttpOnly = true

	SessionStore = stor
	scps := map[string]bool{
		"secure.client.query": true,
	}

	token, err := middle.FetchToken(http.DefaultClient, securityUrl, clnt, scrt, "", scps)

	if err != nil {
		panic(err)
	}

	Authority = kong.NewAuthority(http.DefaultClient, securityUrl, managerUrl, token)

	tmpl, err := drx.LoadTemplate("./views")

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))

	clntIns := middle.NewClientInspector(clnt, scrt, http.DefaultClient, securityUrl, "")
	r.HandleFunc("/", clntIns.Middleware(Index(tmpl), make(map[string]bool))).Methods(http.MethodGet)
	r.HandleFunc("/login", clntIns.Middleware(LoginGET(tmpl), map[string]bool{"entity.login.apply": true})).Queries("state", "{state}", "client", "{client}").Methods(http.MethodGet)

	r.HandleFunc("/consent", clntIns.Middleware(ConsentGET(tmpl), map[string]bool{"entity.consent.apply": true, "secure.client.query": true})).Queries("client", "{client}", "callback", "{callback}", "state", "{state}").Methods(http.MethodGet)
	r.HandleFunc("/consent", clntIns.Middleware(ConsentUserGET(tmpl), map[string]bool{"entity.consent.apply": true, "secure.client.query": true})).Queries("client", "{client}", "state", "{state}", "partial", "{partial}").Methods(http.MethodGet)

	r.HandleFunc("/forgot", clntIns.Middleware(Index(tmpl), make(map[string]bool))).Methods(http.MethodGet)
	r.HandleFunc("/register", clntIns.Middleware(Index(tmpl), make(map[string]bool))).Methods(http.MethodGet)
	return r
}
