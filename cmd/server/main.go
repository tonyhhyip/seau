package main

import (
	"flag"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/tonyhhyip/seau/pkg/server/setup"
)

func init() {
	var verbose bool
	flag.BoolVar(&verbose, "verbose", false, "Verbose for debug")
	flag.Parse()
	if verbose {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Debug("Enable Verbose")
	}
	setup.Bootstrap()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, setup.GetHandler()); err != nil {
		panic(err)
	}
}
