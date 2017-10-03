package handlers

import (
	"encoding/json"
	"github.com/huangjoyce3/info344-in-class/zipsvr/models"
	"net/http"
	"strings"
)

// create handler to return zipcodes of a given city
type CityHandler struct {
	PathPrefix string
	Index      models.ZipIndex
}

// ch is the this pointer
func (ch *CityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// URL: /zips/city-name
	cityName := r.URL.Path[len(ch.PathPrefix):] // slice syntax works on strings
	cityName = strings.ToLower(cityName)
	if len(cityName) == 0 { // error message
		http.Error(w, "please provide a city name", http.StatusBadRequest)
		return
	}

	w.Header().Add(headerContentType, contentTypeJSON) // refer to constants.go
	w.Header().Add(headerAccessControlAllowOrigin, "*")
	zips := ch.Index[cityName]
	json.NewEncoder(w).Encode(zips)
}
