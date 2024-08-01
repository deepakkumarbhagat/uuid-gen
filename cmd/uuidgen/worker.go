package main

import (
	"fmt"
	"net/http"

	"github.com/deepakkumarbhagat/uuidgen"
)

type UUIDS struct {
	Uuid uuidgen.UUID `json:"uuid"`
}

type Worker struct {
	ch         chan Job
	id         WorkerID
	workerPool chan WorkerID
	metaData   uuidgen.MetaData
}

func NewWorker(ch chan Job, id WorkerID, wp chan WorkerID, mData uuidgen.MetaData) *Worker {
	w := &Worker{
		ch:         ch,
		id:         id,
		workerPool: wp,
		metaData:   mData}

	return w
}

func (w *Worker) Start() {
	fmt.Printf("Starting workers with id: %d", int(w.id))
	fmt.Println()
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
	job.Ctx.IndentedJSON(http.StatusOK, UUIDS{Uuid: uuidgen.GenerateUUID(w.metaData)})
}
