package builder

import (
	"context"
	"fmt"

	remoteexecution "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	re_filesystem "github.com/buildbarn/bb-remote-execution/pkg/filesystem"
	"github.com/buildbarn/bb-remote-execution/pkg/proto/remoteworker"
	"github.com/buildbarn/bb-storage/pkg/clock"
	"github.com/buildbarn/bb-storage/pkg/digest"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type timestampedBuildExecutor struct {
	buildExecutor BuildExecutor
	clock         clock.Clock
	workerName    string
}

// NewTimestampedBuildExecutor creates a decorator for BuildExecutor
// that augments the ActionResult that is part of the ExecuteResponse
// with ExecutedActionMetadata. More concretely, it ensures that the
// ActionResult contains the name of the worker performing the build and
// timing information.
func NewTimestampedBuildExecutor(buildExecutor BuildExecutor, clock clock.Clock, workerName string) BuildExecutor {
	return &timestampedBuildExecutor{
		buildExecutor: buildExecutor,
		clock:         clock,
		workerName:    workerName,
	}
}

func (be *timestampedBuildExecutor) getCurrentTime() *timestamp.Timestamp {
	now, err := ptypes.TimestampProto(be.clock.Now())
	if err != nil {
		panic(fmt.Sprintf("Failed to obtain current timestamp: %s", err))
	}
	return now
}

func (be *timestampedBuildExecutor) Execute(ctx context.Context, filePool re_filesystem.FilePool, instanceName digest.InstanceName, request *remoteworker.DesiredState_Executing, executionStateUpdates chan<- *remoteworker.CurrentState_Executing) *remoteexecution.ExecuteResponse {
	// Initial metadata, using the current time as the start timestamp.
	metadata := remoteexecution.ExecutedActionMetadata{
		Worker:               be.workerName,
		QueuedTimestamp:      request.QueuedTimestamp,
		WorkerStartTimestamp: be.getCurrentTime(),
	}

	// Call into the underlying build executor.
	baseUpdates := make(chan *remoteworker.CurrentState_Executing)
	baseCompletion := make(chan *remoteexecution.ExecuteResponse)
	go func() {
		baseCompletion <- be.buildExecutor.Execute(ctx, filePool, instanceName, request, baseUpdates)
	}()

	var completedTimestamp **timestamp.Timestamp
	for {
		select {
		case update := <-baseUpdates:
			// Complete the previous stage.
			now := be.getCurrentTime()
			if completedTimestamp != nil {
				*completedTimestamp = now
			}

			// Start the next stage.
			switch update.ExecutionState.(type) {
			case *remoteworker.CurrentState_Executing_FetchingInputs:
				metadata.InputFetchStartTimestamp = now
				completedTimestamp = &metadata.InputFetchCompletedTimestamp
			case *remoteworker.CurrentState_Executing_Running:
				metadata.ExecutionStartTimestamp = now
				completedTimestamp = &metadata.ExecutionCompletedTimestamp
			case *remoteworker.CurrentState_Executing_UploadingOutputs:
				metadata.OutputUploadStartTimestamp = now
				completedTimestamp = &metadata.OutputUploadCompletedTimestamp
			default:
				completedTimestamp = nil
			}
			executionStateUpdates <- update
		case response := <-baseCompletion:
			// Complete the final stage.
			now := be.getCurrentTime()
			if completedTimestamp != nil {
				*completedTimestamp = now
			}

			// Merge the metadata into the response.
			metadata.WorkerCompletedTimestamp = now
			proto.Merge(response.Result.ExecutionMetadata, &metadata)
			return response
		}
	}
}
