package js

import (
	"go.uber.org/zap"
	v8 "rogchap.com/v8go"
)

// Handles JS VM and creating of new contexts
type JsVM struct {
	isolate *v8.Isolate
	// Logger is included here for ergonomics
	log *zap.Logger
}

// Create a new VM
func New(log *zap.Logger) JsVM {
	log.Info("Creating JS VM")

	return JsVM{v8.NewIsolate(), log}
}

// Create a new context to execute JS in a sandbox environment
func (vm *JsVM) NewContext() *v8.Context {
	vm.log.Info("New JS context created")

	return v8.NewContext(vm.isolate)
}
