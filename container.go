package godi

import "fmt"

type Container struct {
	deps Dependencies
}

/*
   Get dependency by name defined on Builder.

   If dependency not exists, it panics.
   If Create function return error, it panics.
*/
func (ctn Container) Get(name string) interface{} {
	data, exists := ctn.deps[name]
	if !exists {
		panic(fmt.Errorf("no dependencies called %s", name))
	}

	val, err := data.Create()
	if err != nil {
		panic(err)
	}

	return val
}
