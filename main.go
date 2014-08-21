package main

import (
	"flag"
	"github.com/die-net/unifi"
	"log"
	"net/http"
	"runtime"
)

var (
	host    = flag.String("host", "unifi", "Controller hostname")
	user    = flag.String("user", "admin", "Controller username")
	pass    = flag.String("pass", "unifi", "Controller password")
	version = flag.Int("version", 2, "Controller base version")
	siteid  = flag.String("siteid", "default", "Site ID, UniFi v3 only")
	listen  = flag.String("listen", ":8000", "The [IP]:port to listen for incoming connections on.")
	threads = flag.Int("threads", runtime.NumCPU(), "The number of worker threads to execute.")
)

func main() {
	runtime.GOMAXPROCS(*threads)

	flag.Parse()

	u, err := unifi.Login(*user, *pass, *host, *siteid, *version)
	if err != nil {
		log.Fatal(err)
	}
	defer u.Logout()

	log.Fatal(http.ListenAndServe(*listen, nil))
}
