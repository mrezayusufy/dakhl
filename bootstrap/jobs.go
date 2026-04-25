package bootstrap

import (
	"github.com/goravel/framework/contracts/queue"

	"dakhl/app/jobs"
)

func Jobs() []queue.Job {
	return []queue.Job{
		&jobs.Test{},
		&jobs.TestErr{},
	}
}
