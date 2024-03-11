package delayedjob

import "github.com/baselrabia/delayed-job-scheduler/pkg/job"

// delayedJob is the wrapper of Job which contains extra information required for the priority queue implementation
type delayedJob struct {
	job job.Job
	// priority is the epoch time at which this job needs to be executed
	priority int64
	// index is the position of this Job in the priorityQueue
	index int
}

// priorityJobQueue is the priority queue implementation which will store the delayedJob in their order of execution
type priorityJobQueue []*delayedJob

func (pq priorityJobQueue) Len() int { return len(pq) }

func (pq priorityJobQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq priorityJobQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityJobQueue) Push(x interface{}) {
	n := len(*pq)
	delayedJob := x.(*delayedJob)
	delayedJob.index = n
	*pq = append(*pq, delayedJob)
}

func (pq *priorityJobQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	delayedJob := old[n-1]
	old[n-1] = nil        // avoid memory leak
	delayedJob.index = -1 // for safety
	*pq = old[0 : n-1]
	return delayedJob
}
