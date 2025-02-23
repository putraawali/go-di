package godi

type Dependency struct {
	// Name is the key of the dependencies, should be unique for each dependency
	Name string

	// Create is a function to create the object of dependency
	Create func() (interface{}, error)
}

type Dependencies map[string]Dependency
