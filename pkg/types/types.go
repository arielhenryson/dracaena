package types

type Workflow struct {
	Name string
	Jobs map[string] *WorkflowJobRunner
	Context *WorkflowContext
	Internal WorkflowInternal
}

type WorkflowInternal struct {
	Status string
	CompletedJobs int
	ID string
}

type WorkflowContext struct {
	Git string
}

type WorkflowJobRunner struct {
	Runner string
	Payload *WorkFlowPayload
	FailFast bool
	Async bool
	Internal JobInternal
}

type JobInternal struct {
	Status string
	JobsToCompleteBeforeStart int
}

type BusMessage struct {
	JobId string
	Type string
	Payload string
}

type WorkFlowPayload struct {
	Image string
	Commands []string
}