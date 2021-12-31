package api

import (
	"net/http"

	"github.com/everstake/cosmoscan-api/log"
	"github.com/gorilla/mux"
)

func (api *API) GetAccount(w http.ResponseWriter, r *http.Request) {
	address, ok := mux.Vars(r)["address"]
	if !ok || address == "" {
		jsonBadRequest(w, "invalid address")
		return
	}
	resp, err := api.svc.GetAccount(address)
	if err != nil {
		log.Error("API GetAccount: svc.GetAccount: %s", err.Error())
		jsonError(w)
		return
	}
	jsonData(w, resp)
}
