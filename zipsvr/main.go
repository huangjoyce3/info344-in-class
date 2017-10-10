package main

import (
	"encoding/json"
	"fmt"
	"github.com/huangjoyce3/info344-in-class/zipsvr/handlers"
	"github.com/huangjoyce3/info344-in-class/zipsvr/models"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
)

const zipsPath = "/zips/"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello %s!", name) // http://localhost:4000/hello?name=Joyce
}

func memoryHandler(w http.ResponseWriter, r *http.Request) {
	runtime.GC()
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(stats)
}

func main() {
	addr := os.Getenv("ADDR") // environment vars are caps by convention
	if len(addr) == 0 {
		addr = ":80" // in terminal: export ADDR=localhost:4000
	}

	tlskey := os.Getenv("TLSKEY")
	tlscert := os.Getenv("TLSCERT")
	if len(tlskey) == 0 || len(tlscert) == 0 {
		log.Fatal("please set TLSKEY and TLSCERT")
	}

	// load zips.csv
	zips, err := models.LoadZips("zips.csv")
	if err != nil {
		log.Fatalf("error handling zips: %v", err) //do not do this for normal error handling
	}
	log.Printf("loaded %d zipz", len(zips))

	// implementing a map to get all seattle zipcodes fastest
	// this is v efficient
	cityIndex := models.ZipIndex{} // static initializer
	for _, z := range zips {       // _ is the index, z is the item (zip code)
		cityLower := strings.ToLower(z.City)
		cityIndex[cityLower] = append(cityIndex[cityLower], z)
	}

	//fmt.Println("Hello world!")
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/memory", memoryHandler)

	cityHandler := &handlers.CityHandler{
		Index:      cityIndex,
		PathPrefix: zipsPath,
	}
	mux.Handle(zipsPath, cityHandler)

	fmt.Printf("server is listening at http://%s\n", addr) // echo that we are listening at this address
	log.Fatal(http.ListenAndServe(addr, mux))
}
