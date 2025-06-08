// Crow is a Go library designed to create command-line applications in a simple and intuitive way using struct fields and tags. Inspired by projects like Commandeer, Crow aims to provide a more straightforward and "plug & play" solution for creating small applications or scripts, thereby reducing the complexity often associated with libraries like Cobra.
//	type MyCommand struct {
//	    Name string `help:"You Name"`
//	    Age  int    `help:"Your age"`
//	}
//
//	func (mc *MyCommand) Run() error {
//	    // Do your stuff here
//	}
//
//	func main() {
//	    app := crow.New("App Name", "App Description")
//	    command := &MyCommand{
//	        Name: "Lucas",
//	        Age: 27,
//	    }
//	    app.AddCommand(command, "Description of the command")
//	    err = app.Execute(os.Args)
//	    if err != nil {
//	        fmt.Println(err)
//	        os.Exit(1)
//	    }
//	}

package crow

import (
	"fmt"
	"regexp"
)

// App represents a command-line application with custom commands.
type App struct {
	Name                string            // Name of the application
	Description         string            // Description of the application
	Commands            []Command         // List of available commands
	Topics              map[string]*Topic // List of available Topics associated with their names
	CommandsDescription map[string]string // Description of commands associated with their names
	Arguments           []string          // Arguments passed to the application
}

// New is a constructor to create a new instance of Crow App.
// It initializes the Name and Description fields and creates an empty map for command descriptions.
func New(name, description string) *App {
	return &App{
		Name:                name,
		Description:         description,
		Topics:              make(map[string]*Topic),
		CommandsDescription: make(map[string]string),
	}
}

// AddCommand adds a new command to the application.
// It takes a Command and its description as arguments.
// It returns an error if the command already exists.
func (app *App) AddCommand(command Command, description string) error {
	// Get the name of the command
	cmdName, err := getNameOfCommand(command)
	if err != nil {
		return err
	}
	// Check if the command already exists
	if _, ok := app.CommandsDescription[cmdName]; ok {
		return fmt.Errorf("The command %s already exist", cmdName)
	}

	// Add the command and its description
	app.Commands = append(app.Commands, command)
	app.CommandsDescription[cmdName] = description

	return nil
}

// AddTopic adds a new help Topic to the application.
// It takes a name, a short description and a long content as arguments.
// It returns an error if the topic already exists.
func (app *App) AddTopic(name, description, content string) error {
	// Check if the topic already exists
	if _, ok := app.Topics[name]; ok {
		return fmt.Errorf("The command %s already exist", name)
	}

	// Add the topic
	app.Topics[name] = &Topic{Description: description, Content: content}

	return nil
}

// Execute processes the arguments passed to the application and executes the appropriate command.
// If no arguments are provided or if the user asks for help, it calls helpHandler.
// Otherwise, it calls commandsHandler to process the commands.
func (app *App) Execute(args []string) error {

	// Add the arguments to the application's arguments list
	app.Arguments = append(app.Arguments, args...)

	// Use a regular expression to check if the user is asking for help
	help := regexp.MustCompile(helpRegexp)

	// If no arguments are provided or if the user asks for help, call helpHandler
	if len(app.Arguments) <= 1 || help.MatchString(app.Arguments[1]) {
		return app.helpHandler()
	}

	// Otherwise, process the commands
	return app.commandsHandler()
}
