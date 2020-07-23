package setup

import (
	"github.com/navinds25/EviveInDesignServer/pkg/indesign"
	"github.com/navinds25/styx/pkg/execute"
)

// QueueSetup sets up all queues for styx
func QueueSetup() {
	execute.SubscribeCmdAsyncParallel()
	indesign.SubscribeInDesignJobs()
}
