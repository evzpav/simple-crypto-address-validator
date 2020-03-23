package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", index)
	router.GET("/validate/:crypto/:address", validateAddressHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	log.Printf("Server is running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

type validationReturn struct {
	Ok      bool   `json:"ok"`
	Crypto  string `json:"crypto"`
	Address string `json:"address"`
	IsValid bool   `json:"valid"`
	Error   string `json:"error,omitempty"`
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Cryptocurrency address validator!\n")
}

func validateAddressHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	crypto := ps.ByName("crypto")
	if crypto == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	address := ps.ByName("address")

	if address == "" || len(address) < 4 {
		w.WriteHeader(http.StatusBadRequest)
		return
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
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(valReturn)

}

func validateAddress(crypto, address string) (bool, error) {
	crypto = strings.ToLower(crypto)

	var cryptoRegexMap = map[string]string{
		"btc":   "^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$",
		"btg":   "^([GA])[a-zA-HJ-NP-Z0-9]{24,34}$",
		"dash":  "^([X7])[a-zA-Z0-9]{33}$",
		"dgb":   "^(D)[a-zA-Z0-9]{24,33}$",
		"eth":   "^(0x)[a-zA-Z0-9]{40}$",
		"smart": "^(S)[a-zA-Z0-9]{33}$",
		"xrp":   "^(r)[a-zA-Z0-9]{33}$",
		"zcr":   "^(Z)[a-zA-Z0-9]{33}$",
		"zec":   "^(t)[a-zA-Z0-9]{34}$",
	}

	regex, ok := cryptoRegexMap[crypto]
	if !ok {
		return false, fmt.Errorf("Cryptocurrency not available: %s ", crypto)
	}

	re := regexp.MustCompile(regex)

	return re.MatchString(address), nil

}
