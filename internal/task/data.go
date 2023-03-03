package task

import (
	"runtask/internal/js"

	"github.com/tidwall/gjson"
	v8 "rogchap.com/v8go"
)

// The output and input type for all tasks.
// Data is stored in key-value format or in specified type stores.
// Also has references to other resources tasks might need.
type DataStore struct {
	// Any valid JSON data
	json string
	jsVM *js.JsVM
}

// Create a new data store, usually one for each workflow
func New(vm *js.JsVM) DataStore {
	return DataStore{jsVM: vm}
}

// Overwrite the current JSON store content
func (store *DataStore) SaveJson(data string) {
	store.json = data
}

// Search the current JSON store content
func (store *DataStore) SearchJson(query string) gjson.Result {
	return gjson.Get(store.json, query)
}

// Get a new context for running JS
func (store *DataStore) NewJsContext() *v8.Context {
	return store.jsVM.NewContext()
}
