package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/kong/prime"
	"html/template"
	"log"
	"net/http"
)

func ConsentGET(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Consent", tmpl, "./views/consent.html")
	return func(w http.ResponseWriter, r *http.Request) {
		sessn, err := SessionStore.Get(r, "partial")

		if err != nil {
			log.Println("Session Error", err)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		ut, ok := sessn.Values["user.token"]

		if !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		req := prime.QueryRequest{Token: ut.(string)}

		user, concern, err := Authority.ClientQuery(req)

		if err != nil {
			log.Println("Client Query Error", err)
			http.Error(w, "", http.StatusInternalServerError)
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

		result := struct {
			ID       string
			Username string
			Concern  map[string][]string
		}{
			ID:       clnts[0],
			Username: user,
			Concern:  concern,
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

/*
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
}*/

/*


 */
