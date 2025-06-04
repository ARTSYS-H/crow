package crow

import (
	"fmt"
	"regexp"
)

type App struct {
	Name                string
	Description         string
	Commands            []Command
	CommandsDescription map[string]string
	Arguments           []string
}

func New(name, description string) *App {
	return &App{
		Name:                name,
		Description:         description,
		CommandsDescription: make(map[string]string),
	}
}

func (app *App) AddCommand(command Command, description string) error {
	cmdName, err := getNameOfCommand(command)
	if err != nil {
		return err
	}
	if _, ok := app.CommandsDescription[cmdName]; ok {
		return fmt.Errorf("The command %s already exist", cmdName)
	}

	app.Commands = append(app.Commands, command)
	app.CommandsDescription[cmdName] = description

	return nil
}

func (app *App) Execute(args []string) error {

	app.Arguments = append(app.Arguments, args...)

	help := regexp.MustCompile(helpRegexp)

	if len(app.Arguments) <= 1 || help.MatchString(app.Arguments[1]) {
		return app.helpHandler()
	}

	return app.commandsHandler()
}
