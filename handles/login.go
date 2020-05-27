package handles

import (
	"fmt"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/kong/prime"
	"github.com/louisevanderlith/kong/samples/servers/auth"
	"html/template"
	"log"
	"net/http"
)

func LoginGET(mstr *template.Template, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.New(w, r)

		mxr := mix.Page("Login", nil, ctx.GetTokenInfo(), mstr, tmpl)

		ctx.Serve(http.StatusOK, mxr)
	}
}

func LoginPOST(w http.ResponseWriter, r *http.Request) {
	clnts := r.URL.Query()["client"]

	if len(clnts) == 0 {
		http.Error(w, "no client query", http.StatusBadRequest)
		return
	}

	cbUrls := r.URL.Query()["callback"]

	if len(cbUrls) == 0 {
		http.Error(w, "no callback query", http.StatusBadRequest)
		return
	}

	obj := prime.LoginRequest{
		Client:   clnts[0],
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	part, err := auth.Security.AuthenticateUser(obj)

	if err != nil {
		log.Println(err)
		//Show login again
		return
	}

	session, err := auth.SessionStore.Get(r, "partial")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["user.token"] = part

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/consent?client=%s&callback=%s",clnts[0], cbUrls[0]), http.StatusFound)
}