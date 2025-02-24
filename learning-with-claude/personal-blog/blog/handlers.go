package blog

import "net/http"

type NewHomeHandler struct {
	posts []BlogPost
}

func (h *NewHomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
