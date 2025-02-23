# GO-DI

Simple Dependency Injection library for go.

## Installing

```cmd
go get github.com/putraawali/go-di
```

# Basic Usage

## Build Container

For dependencies that always needs to be created before use

```go
builder := godi.New()

err := builder.Add(
		godi.Dependency{
			Name: "MySQL",
			Create: func() (interface{}, error) {
                mysqlCon := "mySQL connection"
				return mysqlCon, nil
			},
		},
		godi.Dependency{
			Name: "Mongo",
			Create: func() (interface{}, error) {
                mongoConn := "mongo connection"
				return mongoConn, nil
			},
		},
	)
if err != nil {
    return
}

ctn := builder.Build()
```

For dependencies that pre-build and do not need to be created for each use

```go
builder := godi.New()

logger := "your logging package"

err := builder.Set("Logger", logger)
if err != nil {
    return
}

ctn := builder.Build()
```

### Warning

-   Name and Create field is required, empty value will create panics.
-   Your name for each dependency should be unique. Duplicate name of dependencies will create panics.

## Using Container

Get the created container defined on Builder

```go
mongoDI := builder.Get("Mongo")
```

### Warning

-   If you want to get undefined dependency on initialization, will create panics.
-   If your object on the container return error when created, will create panics
