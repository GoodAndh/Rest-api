package middlerware

import (
	"encoding/json"
	"net/http"
	"restlast/helper"
	"restlast/model"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}
func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Add("X-API-Key", "KEYWEB")
	w.Header().Add("Content-Type", "application/json")
	if "KEYWEB" == r.Header.Get("X-API-Key") {
		m.Handler.ServeHTTP(w, r)
	} else {

		webResponse := model.MainWeb{
			Code:   http.StatusUnauthorized,
			Status: "error unauthorized",
		}
		encoder := json.NewEncoder(w)
		err := encoder.Encode(webResponse)
		helper.ErrorFatal(err)
	}
	
}
