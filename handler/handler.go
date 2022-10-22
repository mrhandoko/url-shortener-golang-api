package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (h *handler) GenerateURL(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var generated_url entity.URLshortener
	json.Unmarshal(body, &generated_url)

	id, err := h.service.GenerateShortenURL(entity.URLshortener{
		OriginalURL: generated_url.OriginalURL,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"short_url": fmt.Sprintf("%s/%s", entity.URL, id)})
}

func (h *handler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "invalid param", http.StatusBadRequest)
		return
	}

	data, err := h.service.RedirectURL(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, data.OriginalURL, http.StatusFound)
}

func (h *handler) GetURLData(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetURLs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (h *handler) GetURLByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "invalid param", http.StatusBadRequest)
		return
	}

	data, err := h.service.GetURLByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (h *handler) UpdateURLByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id := param["id"]
	body, _ := ioutil.ReadAll(r.Body)
	var shortenUrl entity.URLshortener
	json.Unmarshal(body, &shortenUrl)

	data, err := h.service.UpdateURLByID(id, shortenUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (h *handler) DeleteURLByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "invalid param", http.StatusBadRequest)
		return
	}

	data, err := h.service.DeleteURLByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
