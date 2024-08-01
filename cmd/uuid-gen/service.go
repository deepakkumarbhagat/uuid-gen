package main

import (
	"strconv"

	"github.com/deepakkumarbhagat/uuidgen"
	"github.com/gin-gonic/gin"
)

type Job struct {
	Ctx *gin.Context
}

type UUIDService struct {
	Disp    *Dispatcher
	JobCh   chan Job
	Workers []*Worker
}

func NewUUIDService(metaData uuidgen.MetaData, maxJobs, maxWorkers string) *UUIDService {
	jobsCapacity, _ := strconv.Atoi(maxJobs)
	workers, _ := strconv.Atoi(maxWorkers)

	jobChannel := make(chan Job, jobsCapacity)
	srv := &UUIDService{
		JobCh: jobChannel,
		Disp:  NewDispatcher(jobChannel, make(chan struct{}), workers),
	}

	for wid := 0; wid < workers; wid++ {
		w := NewWorker(srv.Disp.WorkerCh[wid], WorkerID(wid), srv.Disp.WorkerPool, metaData)
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
	srv.Disp.Stop()
}

func (srv *UUIDService) GetUUID(c *gin.Context) {
	srv.JobCh <- Job{Ctx: c}
}
