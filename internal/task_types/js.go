package tasktypes

import "runtask/internal/task"

// Run JavaScript as a task. The DataStore can input data to the JS VM.
type Js struct {
	task.DefinitionProperties
	code        string
	timeout     uint
	saveJsonOut bool
}

func (js Js) Get() func(data *task.DataStore) error {
	return func(data *task.DataStore) error {
		// Create a new context to run the code
		jsContext := data.NewJsContext()

		jsContext.RunScript(js.code, "main.js")

		return nil
	}
}
