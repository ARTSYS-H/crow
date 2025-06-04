package crow

import (
	"flag"
	"fmt"
	"reflect"
	"strings"
)

// Command is an interface that defines the Run method which all commands must implement.
type Command interface {
	Run() error
}

// commandsHandler handles the execution of commands.
// It iterates over the available commands and executes the one that matches the provided argument.
func (app *App) commandsHandler() error {
	// Iterate over available commands
	for _, cmd := range app.Commands {
		// Get the name of the command
		cmdName, err := getNameOfCommand(cmd)
		if err != nil {
			return err
		}
		// Check if the command name matches the provided argument
		if cmdName == app.Arguments[1] {
			// Get the FlagSet for the command
			fs, err := getFlagSet(cmd)
			if err != nil {
				return err
			}
			// Parse the command arguments
			err = fs.Parse(app.Arguments[2:])
			if err != nil {
				return err
			}
			// Execute the command
			return cmd.Run()
		}
	}
	// Return an error if the command is not found
	return fmt.Errorf("%s %s: Unknown command\nRun '%s help' for usage.", app.Name, app.Arguments[1], app.Name)
}

// getFlagSet creates and configures a FlagSet for a given command.
// It uses reflection to inspect the command's fields and set up corresponding flags.
func getFlagSet(cmd Command) (*flag.FlagSet, error) {
	// Get the name of the command
	cmdName, err := getNameOfCommand(cmd)
	if err != nil {
		return nil, err
	}
	// Inspect and access the value of the command
	valueOfCmd, err := inspectAndAccessValueOfCommand(cmd)
	if err != nil {
		return nil, err
	}

	// Create a new FlagSet for the command
	fs := flag.NewFlagSet(cmdName, flag.ExitOnError)

	// Iterate over the command's fields
	for i := range valueOfCmd.NumField() {
		// Get the "help" and "flag" tags of the field
		help, helpExist := valueOfCmd.Type().Field(i).Tag.Lookup("help")
		flg, flgExist := valueOfCmd.Type().Field(i).Tag.Lookup("flag")

		// Check if the field has a "help" or "flag" tag
		if helpExist || flgExist {
			var name string
			// Use the name from the "flag" tag if available, otherwise use the lowercase field name
			if flgExist {
				name = flg
			} else {
				name = strings.ToLower(valueOfCmd.Type().Field(i).Name)
			}
			// Configure the flag based on the field type
			switch valueOfCmd.Type().Field(i).Type.Kind() {
			case reflect.String:
				p := (*string)(valueOfCmd.Field(i).Addr().UnsafePointer())
				fs.StringVar(p, name, valueOfCmd.Field(i).String(), help)
			case reflect.Int:
				p := (*int)(valueOfCmd.Field(i).Addr().UnsafePointer())
				fs.IntVar(p, name, int(valueOfCmd.Field(i).Int()), help)
			case reflect.Bool:
				p := (*bool)(valueOfCmd.Field(i).Addr().UnsafePointer())
				fs.BoolVar(p, name, valueOfCmd.Field(i).Bool(), help)
			case reflect.Float64:
				p := (*float64)(valueOfCmd.Field(i).Addr().UnsafePointer())
				fs.Float64Var(p, name, valueOfCmd.Field(i).Float(), help)
			case reflect.Uint64:
				p := (*uint64)(valueOfCmd.Field(i).Addr().UnsafePointer())
				fs.Uint64Var(p, name, valueOfCmd.Field(i).Uint(), help)
			case reflect.Uint:
				p := (*uint)(valueOfCmd.Field(i).Addr().UnsafePointer())
				fs.UintVar(p, name, uint(valueOfCmd.Field(i).Uint()), help)
			default:
				return nil, fmt.Errorf("encoutered unsupported field type/kind: %#v", valueOfCmd.Field(i))
			}
		}
	}

	return fs, nil
}

// inspectAndAccessValueOfCommand inspects and accesses the value of a command.
// It uses reflection to get the underlying value of the command.
func inspectAndAccessValueOfCommand(cmd Command) (reflect.Value, error) {
	valueOfCmd := reflect.ValueOf(cmd)

	// Traverse interfaces and pointers to get the underlying value
	for valueOfCmd.Kind() == reflect.Interface || valueOfCmd.Kind() == reflect.Pointer {
		valueOfCmd = valueOfCmd.Elem()
	}

	// Check if the underlying value is a struct
	if valueOfCmd.Kind() == reflect.Struct {
		return valueOfCmd, nil
	}

	return reflect.Value{}, fmt.Errorf("Command does not contain a valid struct")
}

// getNameOfCommand gets the name of a command.
// It uses reflection to get the name of the command's type.
func getNameOfCommand(cmd Command) (string, error) {
	// Inspect and access the value of the command
	valueOfCmd, err := inspectAndAccessValueOfCommand(cmd)
	if err != nil {
		return "", err
	}
	// Get the name of the command's type and convert it to lowercase
	nameOfCmd := valueOfCmd.Type().Name()
	nameOfCmd = strings.ToLower(nameOfCmd)
	return nameOfCmd, nil
}
