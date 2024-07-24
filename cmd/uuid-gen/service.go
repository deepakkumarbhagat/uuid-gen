package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UUIDService struct {
	Disp    *Dispatcher
	JobCh   chan Job
	Workers []*Worker
}

func NewUUIDService() *UUIDService {
	jobsCapacity, _ := strconv.Atoi(maxJobs)
	workers, _ := strconv.Atoi(maxWorkers)

	jobChannel := make(chan Job, jobsCapacity)
	srv := &UUIDService{
		JobCh: jobChannel,
		Disp:  NewDispatcher(jobChannel, make(chan struct{})),
	}

	for wid := 0; wid < workers; wid++ {
		w := NewWorker(srv.Disp.WorkerCh[wid], WorkerID(wid), srv.Disp.WorkerPool)
		srv.Workers = append(srv.Workers, w)
	}

	return srv
}

func (srv *UUIDService) Start() {
	//start the dispatcher
	srv.Disp.Start()

	//start the workers
	for _, w := range srv.Workers {
		w.Start()
	}
}

func (srv *UUIDService) Stop() {
	//stop dispatcher
	srv.Disp.Stop()
}

func (srv *UUIDService) GetUUID(c *gin.Context) {
	fmt.Println("Inside GetUUID")
	srv.JobCh <- Job{Ctx: c}
}
