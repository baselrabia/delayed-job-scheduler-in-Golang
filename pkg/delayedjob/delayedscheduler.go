package delayedjob

import (
	"container/heap"
	"fmt"
	"sync"
	"time"

	"github.com/baselrabia/delayed-job-scheduler/pkg/job"
)

type delayedScheduler struct {
	ticker           *time.Ticker
	priorityJobQueue priorityJobQueue
	sync.Mutex
	done chan bool
}

func NewScheduler() *delayedScheduler {
	delayedScheduler := &delayedScheduler{
		priorityJobQueue: make(priorityJobQueue, 0),
		done:             make(chan bool),
		ticker:           time.NewTicker(1 * time.Second),
	}
	heap.Init(&delayedScheduler.priorityJobQueue)
	go delayedScheduler.start()
	return delayedScheduler
}

func (ds *delayedScheduler) start() {
	go func() {
		for {
			select {
			case <-ds.done:
				ds.ticker.Stop()
				return
			case t := <-ds.ticker.C:
				fmt.Println("Tick at ", t)
				ds.Lock()
				if len(ds.priorityJobQueue) != 0 {
					job := heap.Pop(&ds.priorityJobQueue).(*delayedJob)
					now := time.Now().Unix()
					if job.priority-now > 0 {
						heap.Push(&ds.priorityJobQueue, job)
					} else {
						go job.job.Execute()
					}
				}
				ds.Unlock()
			}
		}
	}()
}

func (ds *delayedScheduler) Stop() {
	ds.done <- true
}

func (ds *delayedScheduler) Schedule(job job.Job, duration time.Duration) {
	ds.Lock()
	now := time.Now().Unix()
	timeOfExecution := now + int64(duration.Seconds())
	delayedJob := &delayedJob{
		job:      job,
		priority: timeOfExecution,
	}
	heap.Push(&ds.priorityJobQueue, delayedJob)
	ds.Unlock()
}
