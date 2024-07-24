package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	maxJobs    = os.Getenv("MAX_JOBS")
	maxWorkers = os.Getenv("MAX_WORKERS")
)

const (
	maxJobsDefault    = "100"
	maxWorkersDefault = "10"

/*
configDir  = "/opt/uuid-gen/config"
configFile = "params.conf"
configPath = configDir + "/" + configFile
*/
)

func init() {
	/*
		f, err := os.Open(configPath)
		if err == nil {
			defer f.Close()
		} else {
			log.Println("Unable to open parameter config file %s: %v", configPath, err)
		}
	*/
	if maxJobs == "" {
		maxJobs = maxJobsDefault
	}

	if maxWorkers == "" {
		maxWorkers = maxWorkersDefault
	}
}

func main() {
	srv := NewUUIDService()

	go func() {
		srv.Start()
	}()

	defer srv.Stop()

	router := gin.Default()
	router.GET("/uuids", srv.GetUUID)

	fmt.Println("listening on port 8080")

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal("service couldn't start")
	}
}
