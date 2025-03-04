package handlers

import (
	"albaranes/sqlc/generated"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UsuarioHandler struct {
	queries *generated.Queries
}

func NewUsuarioHandler(queries *generated.Queries) *UsuarioHandler {
	return &UsuarioHandler{queries: queries}
}

func (h *UsuarioHandler) CreateUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario generated.CreateUsuarioParams
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.queries.CreateUsuario(r.Context(), usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UsuarioHandler) GetUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usuario, err := h.queries.GetUsuario(r.Context(), vars["usuario"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(usuario)
}

func (h *UsuarioHandler) ListUsuarios(w http.ResponseWriter, r *http.Request) {
	usuarios, err := h.queries.ListUsuarios(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(usuarios)
}

func (h *UsuarioHandler) UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var usuario generated.UpdateUsuarioParams
	if err := json.NewDecoder(r.Body).Decode(&usuario); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	usuario.Usuario = vars["usuario"]

	err := h.queries.UpdateUsuario(r.Context(), usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *UsuarioHandler) DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	err := h.queries.DeleteUsuario(r.Context(), vars["usuario"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
