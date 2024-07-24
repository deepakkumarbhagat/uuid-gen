package main

import (
	"fmt"
	"net/http"
)

type UUIDS struct {
	Uuid uint64 `json:"uuid"`
}

type Worker struct {
	ch         chan Job
	id         WorkerID
	workerPool chan WorkerID
}

func NewWorker(ch chan Job, id WorkerID, wp chan WorkerID) *Worker {
	w := &Worker{
		ch:         ch,
		id:         id,
		workerPool: wp}

	return w
}

func (w *Worker) Start() {
	go func() {
		for {
			//add itself to the worker pool
			w.workerPool <- w.id

			if job, ok := <-w.ch; ok {
				w.ProcessRequest(job)
			} else {
				//channel closed
				return
			}
		}
	}()
}

func (w *Worker) ProcessRequest(job Job) {
	fmt.Println("Processing request")
	job.Ctx.IndentedJSON(http.StatusOK, UUIDS{Uuid: 1000})
}
