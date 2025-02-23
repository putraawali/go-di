package godi

import (
	"errors"
	"fmt"
)

type Builder struct {
	dependencies Dependencies
}

func New() *Builder {
	return &Builder{
		dependencies: make(Dependencies),
	}
}

/*
Add: adds one or more dependencies to builder.

It returns an error if dependency can not be added such as empty name, duplicate name, or nil Create function
*/
func (b *Builder) Add(deps ...Dependency) error {
	for _, dep := range deps {
		if err := b.adds(dep); err != nil {
			return err
		}
	}

	return nil
}

func (b *Builder) adds(dep Dependency) error {
	if dep.Name == "" {
		return errors.New("name should not empty")
	}

	if dep.Create == nil {
		return errors.New("create should not empty")
	}

	if b.IsExists(dep.Name) {
		return fmt.Errorf("dependency called %s already exists", dep.Name)
	}

	b.dependencies[dep.Name] = dep

	return nil
}

// Set is a shortcut to add a definition for an already built object.
func (b *Builder) Set(name string, obj interface{}) error {
	return b.adds(Dependency{
		Name: name,
		Create: func() (interface{}, error) {
			return obj, nil
		},
	})
}

// IsExists returns true if dependency with the given name already exists
func (b *Builder) IsExists(name string) bool {
	_, ok := b.dependencies[name]
	return ok
}

// Build creates a Container with all dependencies registered in the Builder
func (b *Builder) Build() Container {
	return Container{
		deps: b.dependencies,
	}
}
