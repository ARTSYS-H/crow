package crow

import "fmt"

// helpRegexp is a regular expression pattern to match various forms of help requests.
const helpRegexp = `^-h|--help|-help|help$`

// helpHandler handles the help command logic.
// It checks if a specific command's help is requested or displays general help information.
func (app *App) helpHandler() error {
	// If no specific command is requested, display general help
	if len(app.Arguments) <= 2 {
		helpMessage, err := app.getHelpString()
		if err != nil {
			return err
		}
		fmt.Println(helpMessage)
		return nil
	}

	// Iterate over commands to find a match for the requested help topic
	for _, cmd := range app.Commands {
		cmdName, err := getNameOfCommand(cmd)
		if err != nil {
			return err
		}
		// Check if the command name matches the requested help topic
		if cmdName == app.Arguments[2] {
			// Get the flag set for the command and print its usage
			fs, err := getFlagSet(cmd)
			if err != nil {
				return err
			}
			fs.Usage()
			return nil
		}
	}
	// Return an error if the help topic is not found
	return fmt.Errorf("%s help %s: unknown help topic. Run '%s help'.", app.Name, app.Arguments[2], app.Name)
}

// getHelpString generates a help message string for the application.
// It includes the application description, usage instructions, and a list of available commands.
func (app *App) getHelpString() (string, error) {
	var helpString string

	// Add the application description to the help string
	helpString += fmt.Sprintf("%s\n", app.Description)
	helpString += "\n"
	helpString += "Usage:\n"
	helpString += "\n"
	helpString += fmt.Sprintf("\t%s <command> [arguments]\n", app.Name)
	helpString += "\n"
	// If there are commands, list them with their descriptions
	if len(app.Commands) >= 1 {
		helpString += "The commands are:\n"
		helpString += "\n"
		for _, cmd := range app.Commands {
			cmdName, err := getNameOfCommand(cmd)
			if err != nil {
				return "", err
			}
			// Format each command and its description into the help string
			helpString += fmt.Sprintf("\t%-15s %s\n", cmdName, app.CommandsDescription[cmdName])

		}
		helpString += "\n"
		helpString += fmt.Sprintf("Use \"%s help <command>\" for more information about a command.\n", app.Name)
	}

	return helpString, nil
}
