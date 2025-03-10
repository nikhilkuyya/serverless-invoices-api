package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/nikhilkuyya/invoice-go-app/internal/store"
)

type BankAccountHandler struct {
	bankAccountStore store.BankAccountStore
}

func NewBankAccountHandler(bankAccountStore store.BankAccountStore) *BankAccountHandler {
	return &BankAccountHandler{
		bankAccountStore: bankAccountStore,
	}
}

func (handler *BankAccountHandler) HandleGetBankAccountByID(w http.ResponseWriter, r *http.Request) {
	paramId := chi.URLParam(r,"id")
	if paramId == "" {
		http.NotFound(w, r)
		return
	}

	bankAccountId, err := strconv.ParseInt(paramId,10,64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	bankAccount, err := handler.bankAccountStore.GetBankAccountByID(bankAccountId);

	if err != nil {
		fmt.Println(err)
		http.Error(w,"couldn't get requested data", http.StatusInternalServerError);
		return
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bankAccount)
}

func (handler *BankAccountHandler) HandleCreateBankAccount(w http.ResponseWriter,r *http.Request) {
	var bankAccount store.BankAccount
	err := json.NewDecoder(r.Body).Decode(&bankAccount)
	if err != nil {
		fmt.Println(err)
		http.Error(w,"failed to create Bank Account",http.StatusInternalServerError)
		return
	}
	createdBankAccount, err := handler.bankAccountStore.CreateBankAccount(&bankAccount)

	if err != nil {
		fmt.Println(err)
		http.Error(w,"failed to create bankAccount",http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(createdBankAccount)
}
