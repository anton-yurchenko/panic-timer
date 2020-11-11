package main

import (
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetReportCaller(false)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		DisableTimestamp:       true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	if os.Getenv("TIMEOUT_SECONDS") == "" {
		log.Fatal("missing env.var TIMEOUT_SECONDS")
	}

	t, err := strconv.Atoi(os.Getenv("TIMEOUT_SECONDS"))
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("panic-timer armed! sleeping %v seconds...", t)

	time.Sleep(time.Duration(int64(t)) * time.Second)

	panic("time is up")
}
