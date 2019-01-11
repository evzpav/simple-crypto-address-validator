package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Cryptocurrency address validator!\n")
}

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/validate/:crypto/:address", validateAddressHandler)

	log.Fatal(http.ListenAndServe(":8888", router))
}

type validationReturn struct {
	Ok      bool   `json:"ok"`
	Crypto  string `json:"crypto"`
	Address string `json:"address"`
	IsValid bool   `json:"valid"`
	Error   string `json:"error,omitempty"`
}

func validateAddressHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	crypto := ps.ByName("crypto")
	if crypto == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	address := ps.ByName("address")

	if address == "" || len(address) < 4 {
		w.WriteHeader(http.StatusBadRequest)
	}
	var ret validationReturn
	ret.Ok = true
	isValid, err := validateAddress(crypto, address)
	if err != nil {
		ret.Ok = false
		ret.Error = err.Error()
	}
	ret.IsValid = isValid
	ret.Crypto = crypto
	ret.Address = address
	valReturn, err := json.Marshal(&ret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(valReturn)

}

func validateAddress(crypto, address string) (bool, error) {
	crypto = strings.ToLower(crypto)
	length := len(address)
	oldBtcAddress := (length >= 25 || length <= 34) && (address[0] == '1' || address[0] == '3')
	switch crypto {
	case "btc":
		newBech32Type := (length == 42 || length == 62) && address[0] == 'b'
		return (oldBtcAddress || newBech32Type), nil
	case "bch":
		if strings.Contains(address, ":") {
			address = strings.Split(address, ":")[1]
			return isBchCashAddressValid(address), nil
		}
		return oldBtcAddress || isBchCashAddressValid(address), nil

	case "btg":
		return (length >= 25 || length <= 34) && (address[0] == 'G' || address[0] == 'A'), nil
	case "dgb":
		return (length >= 25 || length <= 34) && (address[0] == 'D'), nil
	case "dsh", "dash":
		return length == 34 && (address[0] == 'X' || address[0] == '7'), nil
	case "eth":
		return length == 42 && address[0:2] == "0x", nil
	case "smart":
		return length == 34 && address[0] == 'S', nil
	case "xrp":
		return length == 34 && address[0] == 'r', nil
	case "zcr":
		return length == 34 && address[0] == 'Z', nil
	case "zec":
		return length == 35 && address[0] == 't', nil
	default:
		return false, errors.New("Cryptocurrency not available: " + crypto)
	}
}

func isBchCashAddressValid(address string) bool {
	return len(address) == 42 && (address[0] == 'p' || address[0] == 'q')
}
