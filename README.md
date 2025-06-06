<div align="center">
    <img src="./assets/images/logo-crow.png" alt="Crow Logo">
    <h1>Crow</h1>
</div>

[![Go Reference](https://pkg.go.dev/badge/github.com/ARTSYS-H/crow/pkg/crow.svg)](https://pkg.go.dev/github.com/ARTSYS-H/crow/pkg/crow)

Crow is a Go library designed to create command-line applications in a simple and intuitive way using struct fields and tags. Inspired by projects like [Commandeer](https://github.com/jaffee/commandeer), Crow aims to provide a more straightforward and "plug & play" solution for creating small applications or scripts, thereby reducing the complexity often associated with libraries like [Cobra](https://github.com/spf13/cobra).

> :warning: **Warning** Crow is still in development. Consider it experimental.

## Project Origin

Crow was developed to meet a specific need: the quick and simple creation of small command-line applications. Personally, I have often found myself creating scripts and small applications where libraries like [Cobra](https://github.com/spf13/cobra) introduced unnecessary complexity. Additionally, tools like [Commandeer](https://github.com/jaffee/commandeer), which share a similar approach, seem to have lost some of their momentum in terms of maintenance. Crow was born to offer a simple and well-maintained alternative.

## Features

- **Simplicity**: Create CLI applications with minimal code and configuration.
- **Use of Structs and Tags**: Define your commands and options directly in Go structures with tags.
- **Automatic Help Message Generation**: Help messages for your commands are automatically generated, making it easier to document and use your application.
- **Plug and Play**: Designed to be easy to integrate and use without complex configuration.
- **Ideal for Small Projects**: Perfect for scripts and small applications where libraries like [Cobra](https://github.com/spf13/cobra) would be excessive.

> :sparkles: **New Feature:** You can create additional help topic.

## Installation

To install Crow, use the `go get` command:

```bash
go get github.com/ARTSYS-H/crow/pkg/crow
```

## Usage

Here is a simple example to get started with Crow:
```go
package main

import (
    "github.com/ARTSYS-H/crow/pkg/crow"
)

type MyCommand struct {
    Name string `help:"You Name"`
    Age  int    `help:"Your age"`
}

func (mc *MyCommand) Run() error {
    // Do your stuff here
}

func main() {
    app := crow.New("App Name", "App Description")
    command := &MyCommand{
        Name: "Lucas",
        Age: 27,
    }
    app.AddCommand(command, "Description of the command")
    err = app.Execute(os.Args)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
```

### Explanation

- **Structs and Tags**: Define your commands and options as fields of a struct with tags to specify command-line options.
- **crow.Execute**: Use this function to parse command-line arguments.

## Contribution

Contributions are welcome! If you have suggestions, bug fixes, or improvements, feel free to open an issue or a pull request.
