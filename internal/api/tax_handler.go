package api

import (
	"net/http"

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

func (teamHandler *TeamHandler) HandleGetTaxes(w http.ResponseWriter, r *http.Request) {
}

func (teamHandler *TeamHandler) HandleGetTaxByID(w http.ResponseWriter, r *http.Request) {}

func (teamHandler *TeamHandler) HandleCreateTax(w http.ResponseWriter, r *http.Request) {}
