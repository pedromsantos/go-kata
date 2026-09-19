package timing

import "time"

// BackgroundJobRunner waits a fixed, arbitrary delay instead of actually
// waiting on the job's completion -- correctness depends on the job
// finishing within the wait window, a race condition disguised as a
// constant. Connascence of Timing.
type BackgroundJobRunner struct {
	jobResult string
}

// StartJob kicks off the job in the background.
func (r *BackgroundJobRunner) StartJob() {
	go func() {
		time.Sleep(300 * time.Millisecond)
		r.jobResult = "done"
	}()
}

// WaitForResult waits a fixed delay, then returns whatever result is
// there -- not necessarily the job's actual result.
func (r *BackgroundJobRunner) WaitForResult() string {
	time.Sleep(1 * time.Second)
	return r.jobResult
}
