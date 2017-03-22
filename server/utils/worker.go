package utils

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/satori/go.uuid"
)

// Work - contains name of the work and the timeout
type Work struct {
	Name     string
	Executor func() (err error)
	Timeout  time.Duration
}

// Worker - base structure for control workers
type Worker struct {
	ID   uuid.UUID
	Work Work

	StopTask chan bool
}

// NewWorker - create and return new worker object
func NewWorker(work Work) Worker {
	// create new worker with unique id
	worker := Worker{
		ID:       uuid.NewV4(),
		Work:     work,
		StopTask: make(chan bool),
	}
	return worker
}

// Start - run the worker
func (w *Worker) Start() {
	go func() {
		log.WithField("uuid", w.ID).Info("The worker with is now running.")
		select {
		case <-w.StopTask:
			log.WithField("uuid", w.ID).Info("The worker has been stopped.")
			return
		}
	}()
}
