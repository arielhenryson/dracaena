// all the types of response runner can send throw the
// orchestrator event bus
package runnerResponseTypes

const (
	// signal runner exit process
	Exit = "exit"

	// signal log from runner
	Log = "log"
)