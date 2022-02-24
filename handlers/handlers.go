package handlers

import (
	"currency-converter/errors"
	"currency-converter/objects"
	"log"
	"net/http"
)

// ICurrencyHandler is implement all the handlers
type ICurrencyHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	cr objects.ICurrency
}

func NewHandler(currencies objects.Currencies) ICurrencyHandler {
	return &handler{cr: &currencies}
}

func (h *handler) Get(w http.ResponseWriter, req *http.Request) {
	amount := req.URL.Query().Get("amount")
	fromCurrency := req.URL.Query().Get("fromCurrency")
	toCurrency := req.URL.Query().Get("toCurrency")

	res, err := h.cr.Convert(amount, fromCurrency, toCurrency)
	if err != nil {
		log.Println(err)
		WriteError(w, errors.ErrBadRequest)
		return
	}
	WriteResponse(w, &objects.ResponseWrapper{Result: res})
}
