package asyncjob

import (
	"context"
	"time"
)

//Job requirement
//1. Job can do something (handler)
//2. Job can retry
// 2.1 I could config retry time and duration
//3. Should be stateful
//4. We should have job manager to manage job

type Job interface {
	Execute(ctx context.Context) error
	Retry(ctx context.Context) error
	State() JobState
	SetRetryDuration(times []time.Duration)
}

const (
	defaultMaxTimeOut    = time.Second * 10
	defaultMaxRetryCount = 3
)

var (
	defaultRetryTime = []time.Duration{time.Second, time.Second * 5, time.Second * 10}
)

type JobHandler func(ctx context.Context) error
type JobState int

const (
	StateInit JobState = iota
	StateRunning
	StateFailed
	StateTimeout
	StateCompleted
	StateRetryFailed
)

func (js JobState) String() string {
	return []string{"Init", "Running", "Failed", "Timeout", "Completed", "RetryFailed"}[js]
}

type jobConfig struct {
	MaxTimeout time.Duration
	Retries    []time.Duration
}

type job struct {
	config     jobConfig
	handler    JobHandler
	state      JobState
	retryIndex int
	stopChan   chan bool
}

func NewJob(handler JobHandler) *job {
	j := job{
		config: jobConfig{
			MaxTimeout: defaultMaxTimeOut,
			Retries:    defaultRetryTime,
		},
		handler:    handler,
		retryIndex: -1,
		state:      StateInit,
		stopChan:   make(chan bool),
	}
	return &j
}

func (j *job) Execute(ctx context.Context) error {
	j.state = StateRunning

	var err error
	err = j.handler(ctx)

	if err != nil {
		j.state = StateFailed
		return err
	}

	j.state = StateCompleted

	return nil
}

func (j *job) Retry(ctx context.Context) error {
	j.retryIndex += 1

	time.Sleep(j.config.Retries[j.retryIndex])

	err := j.Execute(ctx)

	if err == nil {
		j.state = StateCompleted
		return nil
	}

	if j.retryIndex == len(j.config.Retries)-1 {
		j.state = StateRetryFailed
		return nil

	}

	j.state = StateFailed
	return err
}

func (j *job) State() JobState { return j.state }
func (j *job) RetryIndex() int { return j.retryIndex }
func (j *job) SetRetryDuration(times []time.Duration) {
	if len(times) == 0 {
		return
	}

	j.config.Retries = times

}
