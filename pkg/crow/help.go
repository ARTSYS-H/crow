package crow

import "fmt"

const helpRegexp = `^-h|--help|-help|help$`

func (app *App) helpHandler() error {
	if len(app.Arguments) <= 2 {
		helpMessage, err := app.getHelpString()
		if err != nil {
			return err
		}
		fmt.Println(helpMessage)
		return nil
	}

	for _, cmd := range app.Commands {
		cmdName, err := getNameOfCommand(cmd)
		if err != nil {
			return err
		}
		if cmdName == app.Arguments[2] {
			fs, err := getFlagSet(cmd)
			if err != nil {
				return err
			}
			fs.Usage()
			return nil
		}
	}
	return fmt.Errorf("%s help %s: unknown help topic. Run '%s help'.", app.Name, app.Arguments[2], app.Name)
}

func (app *App) getHelpString() (string, error) {
	var helpString string

	helpString += fmt.Sprintf("%s\n", app.Description)
	helpString += "\n"
	helpString += "Usage:\n"
	helpString += "\n"
	helpString += fmt.Sprintf("\t%s <command> [arguments]\n", app.Name)
	helpString += "\n"
	if len(app.Commands) >= 1 {
		helpString += "The commands are:\n"
		helpString += "\n"
		for _, cmd := range app.Commands {
			cmdName, err := getNameOfCommand(cmd)
			if err != nil {
				return "", err
			}
			helpString += fmt.Sprintf("\t%-15s %s\n", cmdName, app.CommandsDescription[cmdName])

		}
		helpString += "\n"
		helpString += fmt.Sprintf("Use \"%s help <command>\" for more information about a command.\n", app.Name)
	}

	return helpString, nil
}
