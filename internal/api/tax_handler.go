package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nikhilkuyya/invoice-go-app/internal/store"
)

type TaxHandler struct {
	taxStore store.TaxStore
}

func NewTaxHandler(taxStore store.TaxStore) *TaxHandler {
		return &TaxHandler{
			taxStore: taxStore,
		}
}

func (taxHandler *TaxHandler) HandleGetTaxes(w http.ResponseWriter, r *http.Request) {
	teams, err := taxHandler.taxStore.GetTaxes()
	if err != nil {
		fmt.Println(err)
		http.Error(w,"Failed to get teams",http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(teams)
}

func (taxHandler *TaxHandler) HandleGetTaxByID(w http.ResponseWriter, r *http.Request) {
	requesteId := chi.URLParam(r,"id");
	if requesteId == "" {
		http.Error(w,"not able to fetch tax information",http.StatusBadRequest)
		return;
	}

	taxId, err := strconv.ParseInt(requesteId,10,64);
	if err != nil {
		fmt.Println(err)
		http.Error(w, "not able to send tax information",http.StatusBadRequest)
		return;
	}

	tax, err := taxHandler.taxStore.GetTaxByID(taxId)

	if err != nil {
		http.Error(w,"Failed to fetch tax information",http.StatusInternalServerError)
		return;
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(tax)

}

func (taxHandler *TaxHandler) HandleCreateTax(w http.ResponseWriter, r *http.Request) {
	var taxPayload store.Tax
	 err :=json.NewDecoder(r.Body).Decode(&taxPayload)

	if err != nil {
		fmt.Println(err)
		http.Error(w,"failed to create tax",http.StatusBadRequest)
		return
	}

	createdTax, err := taxHandler.taxStore.CreateTax(&taxPayload)

	if err != nil {
		fmt.Println(err)
		http.Error(w,"failed to create tax",http.StatusInternalServerError);
		return
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(createdTax)
}
