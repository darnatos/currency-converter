package main

import (
	"currency-converter/handlers"
	"currency-converter/objects"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Args used to run the server
type Args struct {
	fileName string
	// port for the server of the form,
	// e.g ":8000"
	port string
}

// Run the server based on given args
func Run(args Args) error {
	// router

	currencies, err := parseJsonFile(args.fileName)
	if err != nil {
		return err
	}
	router := mux.NewRouter()
	hnd := handlers.NewHandler(currencies)
	RegisterAllRoutes(router, hnd)

	// start server
	log.Println("Listing for requests at http://localhost" + args.port + "/currency")
	return http.ListenAndServe(args.port, router)
}

func RegisterAllRoutes(router *mux.Router, hnd handlers.ICurrencyHandler) {
	// set content type
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	router.HandleFunc("/currency", hnd.Get).Methods(http.MethodGet)
}

func parseJsonFile(fileName string) (currencies objects.Currencies, err error) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		return currencies, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return currencies, err
	}
	json.Unmarshal(byteValue, &currencies)
	return currencies, nil
}
