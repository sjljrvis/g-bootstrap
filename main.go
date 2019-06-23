package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/sjljrvis/g-bootstrap/router"
	"github.com/sjljrvis/g-bootstrap/tools"
	"github.com/subosito/gotenv"
)

const (
	//TimeFormat *generic time stamp foramt used in logger
	TimeFormat = "Jan 2, 2006 at 3:04pm (MST)"
)

func initLogger() {
	log.SetFlags(0)
	log.SetPrefix("app | " + time.Now().Format(TimeFormat) + "| INFO : ")
}

func init() {
	gotenv.Load()
	initLogger()
	tools.ConnectRedis()
}

func main() {
	log.Println("Starting server", os.Getenv("PORT"))

	r := router.NewRouter()
	corsObj := handlers.AllowedOrigins([]string{"*"})
	http.Handle("/", handlers.CORS(corsObj)(r))
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
