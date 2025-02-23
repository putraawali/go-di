package godi_test

import (
	"testing"

	godi "github.com/putraawali/go-di"
	"github.com/stretchr/testify/assert"
)

func TestSuccessBuildDI(t *testing.T) {
	builder := godi.New()

	err := builder.Add(
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

	assert.Nil(t, err)

	ctn := builder.Build()
	assert.NotEmpty(t, ctn)
}

func TestFailedDuplicateKey(t *testing.T) {
	builder := godi.New()

	err := builder.Add(
		godi.Dependency{
			Name: "Repository",
			Create: func() (interface{}, error) {
				return "repository", nil
			},
		},
		godi.Dependency{
			Name: "Repository",
			Create: func() (interface{}, error) {
				return "service", nil
			},
		},
	)

	assert.Error(t, err)
}

func TestFailedEmptyName(t *testing.T) {
	builder := godi.New()

	err := builder.Add(
		godi.Dependency{
			Name: "",
			Create: func() (interface{}, error) {
				return "repository", nil
			},
		},
	)

	assert.Error(t, err)
}

func TestFailedNillCreate(t *testing.T) {
	builder := godi.New()

	err := builder.Add(
		godi.Dependency{
			Name: "Repository",
		},
	)

	assert.Error(t, err)
}

func TestSuccessSet(t *testing.T) {
	builder := godi.New()

	err := builder.Set("Repository", "Repository Layer")

	assert.Nil(t, err)

	ctn := builder.Build()
	assert.NotEmpty(t, ctn)
}

func TestFailedSetEmptyName(t *testing.T) {
	builder := godi.New()

	err := builder.Set("", "Repository Layer")

	assert.Error(t, err)
}

func TestFailedSetDuplicateName(t *testing.T) {
	builder := godi.New()

	builder.Add(godi.Dependency{
		Name: "Repository",
		Create: func() (interface{}, error) {
			return "repository", nil
		},
	})

	err := builder.Set("Repository", "Repository Layer")

	assert.Error(t, err)
}
