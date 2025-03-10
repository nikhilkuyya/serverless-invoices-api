package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nikhilkuyya/invoice-go-app/internal/store"
)


type ClientHandler struct {
	clientStore store.ClientStore
}

func NewClientHandler(clientStore store.ClientStore) *ClientHandler{
	return &ClientHandler{
		clientStore: clientStore,
	}
}

func (handler *ClientHandler) HandleGetClients(w http.ResponseWriter, r *http.Request) {
	clients, err := handler.clientStore.GetClients();

	if err != nil {
		fmt.Println(err)
		http.Error(w,"failed to get clients",http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(clients)
}

func (handler *ClientHandler) HandleCreateClient(w http.ResponseWriter, r *http.Request){
	var clientPayload store.Client
	err := json.NewDecoder(r.Body).Decode(&clientPayload)

	if err != nil {
		fmt.Println(err)
		http.Error(w,"failed to create client",http.StatusInternalServerError)
		return
	}

	 createdClient, err := handler.clientStore.CreateClient(&clientPayload)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to create client",http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(createdClient)
}

func (handler *ClientHandler) HandleGetClientByID(w http.ResponseWriter, r *http.Request) {
	paramId := chi.URLParam(r,"id")

	if paramId == "" {
		http.NotFound(w,r)
		return
	}

	clientId, err := strconv.ParseInt(paramId,10,64)
	if err != nil {
		http.NotFound(w,r)
		return
	}

	client, err := handler.clientStore.GetClientByID(clientId)

	if err != nil {
		fmt.Println(err)
		http.Error(w,"couldn't get the client",http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}
