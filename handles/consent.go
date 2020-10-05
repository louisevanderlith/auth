package handles

import (
	"fmt"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"
)

func ConsentGET(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Consent", tmpl, "./views/consent.html")
	pge.AddMenu(FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {
		cbUrl := drx.FindQueryParam(r, "callback")

		if len(cbUrl) == 0 {
			http.Error(w, "no callback query", http.StatusBadRequest)
			return
		}

		state := drx.FindQueryParam(r, "state")

		if len(state) == 0 {
			log.Println("no 'state' query")
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		client := drx.FindQueryParam(r, "client")

		if len(client) == 0 {
			log.Println("no 'client' query")
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		sessn, err := SessionStore.Get(r, "partial")

		if err != nil {
			log.Println("New Session Error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		sessn.Values[state] = cbUrl

		err = sessn.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		loginUrl := fmt.Sprintf("/login?state=%s&client=%s", state, client)
		http.Redirect(w, r, loginUrl, http.StatusFound)
	}
}

func ConsentUserGET(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Consent", tmpl, "./views/consent.html")
	pge.AddMenu(FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {
		partial := drx.FindQueryParam(r, "partial")

		if len(partial) == 0 {
			log.Println("no 'partial' query")
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		state := drx.FindQueryParam(r, "state")

		if len(state) == 0 {
			log.Println("no 'state' query")
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		client := drx.FindQueryParam(r, "client")

		if len(client) == 0 {
			log.Println("no 'client' query")
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		cc, err := Authority.ClientQuery(client)

		if err != nil {
			log.Println("Client Query Error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		sessn, err := SessionStore.Get(r, "partial")

		if err != nil {
			log.Println("Invalid Session Error", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		cbUrl := sessn.Values[state]

		if len(cbUrl.(string)) == 0 {
			http.Error(w, "no callback found", http.StatusBadRequest)
			return
		}

		result := struct {
			ID       string
			Username string
			Callback string
			Concern  map[string][]string
		}{
			ID:       cc.Client,
			Username: "Userx",
			Callback: cbUrl.(string),
			Concern:  cc.Needs,
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
