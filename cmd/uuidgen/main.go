package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/deepakkumarbhagat/uuidgen"
	"github.com/gin-gonic/gin"
)

var (
	maxJobs      = os.Getenv("MAX_JOBS")
	maxWorkers   = os.Getenv("MAX_WORKERS")
	dataCenterID = os.Getenv("DATACENTER_ID")
	hostID       = os.Getenv("HOST_ID")
)

const (
	maxJobsDefault      = "100"
	maxWorkersDefault   = "10"
	dataCenterIDDefault = "31" // 1<<5 -1
	hostIDDefault       = "31" // 1<<5 -1
)

func init() {
	if maxJobs == "" {
		maxJobs = maxJobsDefault
	}
	if maxWorkers == "" {
		maxWorkers = maxWorkersDefault
	}
	if dataCenterID == "" {
		dataCenterID = dataCenterIDDefault
	}
	if hostID == "" {
		hostID = hostIDDefault
	}
}

func main() {
	dcID, _ := strconv.Atoi(maxJobs)
	hID, _ := strconv.Atoi(maxWorkers)
	metaData := uuidgen.MetaData{
		DataCenterID: dcID,
		HostID:       hID,
	}

	srv := NewUUIDService(metaData, maxJobs, maxWorkers)

	go func() {
		srv.Start()
	}()

	defer srv.Stop()

	router := gin.Default()
	router.GET("/apis/v1/uuids", srv.GetUUID)

	fmt.Println("listening on port 8080")

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatal("service couldn't start")
	}
}
