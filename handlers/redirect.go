package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yashK-2703/go-url-shortener/storage"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	url, err := storage.GetURL(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
