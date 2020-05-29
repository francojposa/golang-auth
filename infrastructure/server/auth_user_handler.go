package server

import (
	"encoding/json"
	"golang-auth/usecases/repos"
	"log"
	"net/http"
)

type httpUserAuthentication struct {
	Username string
	Password string
}

type AuthNUserHandler struct {
	repo repos.AuthNUserRepo
}

func NewAuthNUserHandler(repo repos.AuthNUserRepo) *AuthNUserHandler {
	return &AuthNUserHandler{repo: repo}
}

func (h *AuthNUserHandler) Authenticate(w http.ResponseWriter, r *http.Request) {
	postedUserAuth := httpUserAuthentication{}
	err := json.NewDecoder(r.Body).Decode(&postedUserAuth)
	if err != nil || postedUserAuth.Username == "" || postedUserAuth.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	verified, err := h.repo.Verify(postedUserAuth.Username, postedUserAuth.Password)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !verified {
		w.WriteHeader(http.StatusUnauthorized)
		errBody := map[string]string{"error": "Username or password is incorrect"}
		json.NewEncoder(w).Encode(errBody)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}

}