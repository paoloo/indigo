package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var d = new(DataBase)
var st = new(State)
var t = new(Rtimer)
var m = http.NewServeMux()
var s = http.Server{Addr: ":8080", Handler: m}

func startup(dbname string, rpttime int) {
	log.Printf("[+] starting database %s\n", dbname)
	d.Init(dbname) // create a new json database
	log.Printf("[+] starting memory state.\n")
	st.Init("") // initialize with a hash of an empty memory
	commit := func() { st.CallIfChanged(d.Hash(), d.Store) }
	log.Printf("[+] starting saving timer.\n")
	t.Init(rpttime, commit) // every rpttime ms call d.Store() if the data hash has changed since the last execution
	log.Printf("[+] starting the webserver.\n")
	m.HandleFunc("/", handleRequests)
	log.Fatal(s.ListenAndServe())
}

func cleanup() {
	log.Printf("\n")
	log.Printf("[+] stoping timer.\n")
	t.Stop()
	log.Printf("[+] stoping webserver.\n")
	s.Shutdown(context.Background())
	log.Printf("[+] done.\n")
	os.Exit(1)
}

func main() {
	log.Printf("[+] booting...\n")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	log.Printf("[+] setting the cleanup routine.\n")
	go func() {
		<-c
		cleanup()
	}()
	log.Printf("[+] creating object.\n")
	startup("indigo", 10000)
}
