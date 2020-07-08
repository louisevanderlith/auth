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

func ConsentGET(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage(tmpl, "Consent", "./views/consent.html")
	return func(w http.ResponseWriter, r *http.Request) {
		sessn, err := SessionStore.Get(r, "partial")

		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		ut, ok := sessn.Values["user.token"]

		if !ok {
			log.Println(err)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		req := prime.QueryRequest{Partial: ut.(string)}

		user, concern, err := Security.ClientQuery(req)

		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(nil)
			return
		}

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

		ctx := context.New(w, r)
		result := struct {
			ID       string
			Username string
			Concern  map[string][]string
		}{
			ID:       clnts[0],
			Username: user,
			Concern:  concern,
		}

		mxr := pge.Page(result, ctx.GetTokenInfo(), ctx.GetToken())

		err = ctx.Serve(http.StatusOK, mxr)

		if err != nil {
			log.Println(err)
		}
	}
}

func ConsentPOST(w http.ResponseWriter, r *http.Request) {
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

	session, err := SessionStore.Get(r, "partial")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ut := session.Values["user.token"]

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	obj := prime.ConsentRequest{
		User:   ut.(string),
		Claims: r.Form["consent"],
	}

	tkn, err := auth.Security.GiveConsent(obj)

	if err != nil {
		log.Println(err)
		//Show consent again
		return
	}

	http.Redirect(w, r, fmt.Sprintf("%s?ut=%s", cbUrls[0], tkn), http.StatusFound)
}
