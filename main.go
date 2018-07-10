package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

type Options struct {
	ShowHelp    bool
	Source      string
	Destination string
}

func main() {
	o := parseOptions()

	handler := RedirectHandler{
		Destination: o.Destination,
	}

	listen := fmt.Sprintf(":%s", o.Source)
	fmt.Println("Listening on", listen)
	log.Fatal(http.ListenAndServe(listen, handler))
}

type RedirectHandler struct {
	Destination string
}

func (h RedirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host, _, err := net.SplitHostPort(r.Host)
	redirCode := http.StatusMovedPermanently
	if err != nil {
		host = r.Host
	}
	if r.Method != "GET" {
		// RFC-7538: Redirecting non-GET requests requires a 308 instead of a 301.
		// Otherwise they'll be re-transmitted as GETs, which may break some clients.
		redirCode = http.StatusPermanentRedirect
	}

	destination := *r.URL
	destination.Scheme = "https"
	if h.Destination != "443" {
		destination.Host = net.JoinHostPort(host, h.Destination)
	} else {
		destination.Host = host
	}

	log.Printf(
		"Redirecting to %s %s\n",
		r.Method, destination.String(),
	)

	http.Redirect(w, r, destination.String(), redirCode)
}

func parseOptions() Options {
	var o Options

	flag.BoolVar(&o.ShowHelp, "help", false, "show this help")
	flag.StringVar(&o.Source, "port", "80", "port to listen on")
	flag.StringVar(&o.Destination, "destination", "443", "port to redirect to")

	flag.Parse()

	printHelp := func() {
		fmt.Println("redirect [options]")
		flag.PrintDefaults()
	}

	if o.ShowHelp {
		printHelp()
		os.Exit(0)
	}

	return o
}
