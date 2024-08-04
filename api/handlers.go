package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/luqxus/spaces/service"
	"github.com/luqxus/spaces/types"
)

type handler struct {
	userService *service.UserService
}

func newHandler() *handler {
	return &handler{}
}

func (h *handler) Register(w http.ResponseWriter, r *http.Request) error {
	var reqData *types.RegisterReqData

	if err := DecodeBody(r.Body, reqData); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return nil
	}

	token, err := h.userService.CreateUser(r.Context(), reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	w.Header().Set("authorization", token)

	return WriteJSON(w, http.StatusOK, map[string]string{"message": "user created"})
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) error {
	var reqData *types.LoginReqData

	if err := DecodeBody(r.Body, reqData); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return nil
	}

	user, token, err := h.userService.Login(r.Context(), reqData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	w.Header().Set("authorization", token)

	return WriteJSON(w, http.StatusOK, user)
}

func (api *handler) CreateSpace(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (api *handler) GetSpace(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func DecodeBody(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}

func WriteJSON(w http.ResponseWriter, statusCode int, v any) error {
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(v)
}
