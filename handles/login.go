package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"
)

func LoginGET(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Login", tmpl, "./views/login.html")
	return func(w http.ResponseWriter, r *http.Request) {
		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

/*
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

	http.Redirect(w, r, fmt.Sprintf("/consent?client=%s&callback=%s", clnts[0], cbUrls[0]), http.StatusFound)
}
*/
