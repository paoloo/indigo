package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type infoMsg struct {
	Type        string `json:"type"`
	Information string `json:"information"`
}

func ret2web(w http.ResponseWriter, r *http.Request, itype string, info string) {
	einfo := infoMsg{Type: itype, Information: info}
	bJSON, err := json.MarshalIndent(einfo, "", " ")
	if err != nil {
		log.Println("error:", err)
	}
	// ======================================== CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, STORE, LOAD, DUMP")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	// ================================= set as json
	w.Header().Set("Content-type", "application/json")
	// ============================= return the data
	w.Write(bJSON)
}

func handleRequests(w http.ResponseWriter, r *http.Request) {
	parms := strings.Split(r.URL.Path[1:], "/")
	if r.Method == "GET" && len(parms) > 1 && len(parms[0]) > 0 && len(parms[1]) > 0 {
		ret2web(w, r, "result", d.Get(parms[0], parms[1]))
	} else if r.Method == "POST" && len(parms) > 2 && len(parms[0]) > 0 && len(parms[1]) > 0 && len(parms[2]) > 0 {
		d.Set(parms[0], parms[1], parms[2])
		ret2web(w, r, "result", "ok")
	} else if r.Method == "DELETE" && len(parms) > 1 && len(parms[0]) > 0 && len(parms[1]) > 0 {
		d.Del(parms[0], parms[1])
		ret2web(w, r, "result", "ok")
	} else if r.Method == "STORE" {
		d.Store()
		ret2web(w, r, "result", "ok")
	} else if r.Method == "LOAD" {
		d.Load()
		ret2web(w, r, "result", "ok")
	} else if r.Method == "DUMP" {
		w.Header().Set("Content-type", "application/json")
		w.Write([]byte(d.Dump()))
	} else if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		ret2web(w, r, "error", "Not Implemented yet")
	}
}
