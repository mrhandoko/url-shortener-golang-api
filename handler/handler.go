package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener-golang-api/entity"
	"url-shortener-golang-api/service"

	"github.com/gorilla/mux"
)

type handler struct {
	service service.ServiceInterface
}

func NewHandler(service service.ServiceInterface) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := params["url"]

	id, err := h.service.GenerateShortenURL(entity.URLshortener{
		OriginalURL: url,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"short_url": fmt.Sprintf("%s/%s", entity.URL, id)})
}

func (h *handler) GetStats(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "invalid parameter", http.StatusBadRequest)
		return
	}
	data, err := h.service.GetDataByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (h *handler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "invalid parameter", http.StatusBadRequest)
		return
	}
	data, err := h.service.Redirect(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	http.Redirect(w, r, data.OriginalURL, http.StatusFound)
}
