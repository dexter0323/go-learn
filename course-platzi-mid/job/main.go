package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorker  int
	JobQueue   chan Job
}

func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool),
	}

}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				fmt.Printf("WorkerID: %d started, JobName %v, Fibonacci to calculate %d\n", w.Id, job.Name, job.Number)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("WorkerID: %d started, JobName %v, Fibonacci result: %d\n", w.Id, job.Name, fib)
			case <-w.QuitChan:
				fmt.Printf("WorkerID: %d stopped", w.Id)
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	return &Dispatcher{
		JobQueue:   jobQueue,
		MaxWorker:  maxWorkers,
		WorkerPool: make(chan chan Job, maxWorkers),
	}
}

func (d Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			go func() {
				workerJobQueue := <-d.WorkerPool
				workerJobQueue <- job
			}()

		}
	}
}

func (d Dispatcher) Run() {
	for i := 0; i < d.MaxWorker; i++ {
		w := NewWorker(i, d.WorkerPool)
		w.Start()
	}
	go d.Dispatch()
}

// curl -d "name=HappyFib&value=50&delay=3s" -H "Content-Type: application/x-www-form-urlencoded" -X POST http://localhost:8081/fib
func RequestHandler(w http.ResponseWriter, r *http.Request, jq chan Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid delay", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Invalid name", http.StatusBadRequest)
		return
	}

	j := Job{Name: name, Delay: delay, Number: value}

	jq <- j

}

func main() {
	const (
		maxWorkers   = 4
		maxQueueSize = 20
		port         = ":8081"
	)

	jq := make(chan Job, maxQueueSize)
	d := NewDispatcher(jq, maxQueueSize)
	d.Run()
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jq)
	})
	log.Fatal(http.ListenAndServe(port, nil))
}
