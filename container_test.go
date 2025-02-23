package godi_test

import (
	"errors"
	"testing"

	godi "github.com/putraawali/go-di"
	"github.com/stretchr/testify/assert"
)

func TestSuccessGetContainer(t *testing.T) {
	builder := godi.New()

	builder.Add(
		godi.Dependency{
			Name: "Repository",
			Create: func() (interface{}, error) {
				return "repository", nil
			},
		},
		godi.Dependency{
			Name: "Service",
			Create: func() (interface{}, error) {
				return "service", nil
			},
		},
	)

	ctn := builder.Build()

	repo := ctn.Get("Repository")
	assert.NotNil(t, repo)
	assert.IsType(t, "string", repo)
}

func TestFailedGetContainerNotExists(t *testing.T) {
	defer func() {
		p := recover()

		assert.NotNil(t, p)
	}()

	builder := godi.New()

	builder.Add(
		godi.Dependency{
			Name: "Repository",
			Create: func() (interface{}, error) {
				return "repository", nil
			},
		},
	)

	ctn := builder.Build()

	ctn.Get("repository")
}

func TestFailedGetContainerErrorCreated(t *testing.T) {
	defer func() {
		p := recover()

		assert.NotNil(t, p)
	}()

	builder := godi.New()

	builder.Add(
		godi.Dependency{
			Name: "Repository",
			Create: func() (interface{}, error) {
				return "repository", errors.New("mock error")
			},
		},
	)

	ctn := builder.Build()

	ctn.Get("Repository")
}
