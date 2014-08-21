package main

import (
	"database/sql"
	"flag"
	"github.com/die-net/unifi"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"runtime"
)

var (
	host              = flag.String("host", "unifi", "Controller hostname")
	user              = flag.String("user", "admin", "Controller username")
	pass              = flag.String("pass", "unifi", "Controller password")
	version           = flag.Int("version", 2, "Controller base version")
	siteid            = flag.String("siteid", "default", "Site ID, UniFi v3 only")
	listen            = flag.String("listen", ":8000", "The [IP]:port to listen for incoming connections on.")
	threads           = flag.Int("threads", runtime.NumCPU(), "The number of worker threads to execute.")
	mysqlDsn          = flag.String("mysql_dsn", "utp:@/utp?timeout=2s&parseTime=true&wait_timeout=300", "MySQL connection DSN")
	mysqlMaxIdleConns = flag.Int("mysql_max_idle_conns", 5, "Maximum number of idle MySQL connections to allow.")
	db                *sql.DB
	ufi               *unifi.Unifi
)

func main() {
	flag.Parse()

	runtime.GOMAXPROCS(*threads)

	var err error
	db, err = sql.Open("mysql", *mysqlDsn)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(*mysqlMaxIdleConns)
	defer db.Close()

	ufi, err = unifi.Login(*user, *pass, *host, *siteid, *version)
	if err != nil {
		log.Fatal(err)
	}
	defer ufi.Logout()

	log.Fatal(http.ListenAndServe(*listen, nil))
}
