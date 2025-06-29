package crow

import (
	"fmt"
	"strings"
	"text/tabwriter"
)

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

	// Find a match for the requested Additional help topic
	if topic, exist := app.Topics[app.Arguments[2]]; exist {
		fmt.Println(topic)
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

func (app *App) helpCommandsBuilder(builder *strings.Builder) error {
	if len(app.Commands) >= 1 {
		tw := tabwriter.NewWriter(builder, 8, 0, 3, ' ', 0)
		builder.WriteString("The commands are:\n\n")
		for _, cmd := range app.Commands {
			cmdName, err := getNameOfCommand(cmd)
			if err != nil {
				return err
			}
			// Format each command and its description into the help string
			// fmt.Fprintf(builder, "\t%-15s %s\n", cmdName, app.CommandsDescription[cmdName])
			fmt.Fprintf(tw, "\t%s\t%s\n", cmdName, app.CommandsDescription[cmdName])
		}
		tw.Flush()
		fmt.Fprintf(builder, "\nUse \"%s help <command>\" for more information about a command.\n", app.Name)
	}
	return nil
}

func (app *App) helpTopicsBuilder(builder *strings.Builder) {
	if len(app.Topics) >= 1 {
		tw := tabwriter.NewWriter(builder, 8, 0, 3, ' ', 0)
		builder.WriteString("\nAdditional help topics:\n\n")
		for name, topic := range app.Topics {
			// fmt.Fprintf(builder, "\t%-15s %s\n", name, topic.Short)
			fmt.Fprintf(tw, "\t%s\t%s\n", name, topic.Short)
		}
		tw.Flush()
		fmt.Fprintf(builder, "\nUse \"%s help <topic>\" for more information about a topic.\n", app.Name)
	}
}

// getHelpString generates a help message string for the application.
// It includes the application description, usage instructions, and a list of available commands.
func (app *App) getHelpString() (string, error) {
	var helpBuilder strings.Builder

	// Add the application description to the help string
	fmt.Fprintf(&helpBuilder, "%s\n", app.Description)
	helpBuilder.WriteString("\nUsage:\n\n")
	fmt.Fprintf(&helpBuilder, "\t%s <command> [arguments]\n\n", app.Name)

	// If there are commands, list them with their descriptions
	if err := app.helpCommandsBuilder(&helpBuilder); err != nil {
		return "", err
	}

	// If there are topics, list them with their descriptions
	app.helpTopicsBuilder(&helpBuilder)

	return helpBuilder.String(), nil
}
