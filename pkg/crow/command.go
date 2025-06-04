package crow

import (
	"flag"
	"fmt"
	"reflect"
	"strings"
)

type Command interface {
	Run() error
}

func (app *App) commandsHandler() error {
	for _, cmd := range app.Commands {
		cmdName, err := getNameOfCommand(cmd)
		if err != nil {
			return err
		}
		if cmdName == app.Arguments[1] {
			fs, err := getFlagSet(cmd)
			if err != nil {
				return err
			}
			err = fs.Parse(app.Arguments[2:])
			if err != nil {
				return err
			}
			return cmd.Run()
		}
	}
	return fmt.Errorf("%s %s: Unknown command\nRun '%s help' for usage.", app.Name, app.Arguments[1], app.Name)
}

func getFlagSet(cmd Command) (*flag.FlagSet, error) {
	cmdName, err := getNameOfCommand(cmd)
	if err != nil {
		return nil, err
	}
	valueOfCmd, err := inspectAndAccessValueOfCommand(cmd)
	if err != nil {
		return nil, err
	}

	fs := flag.NewFlagSet(cmdName, flag.ExitOnError)

	for i := range valueOfCmd.NumField() {
		help, helpExist := valueOfCmd.Type().Field(i).Tag.Lookup("help")
		flg, flgExist := valueOfCmd.Type().Field(i).Tag.Lookup("flag")

		if helpExist || flgExist {
			var name string
			if flgExist {
				name = flg
			} else {
				name = strings.ToLower(valueOfCmd.Type().Field(i).Name)
			}
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

func inspectAndAccessValueOfCommand(cmd Command) (reflect.Value, error) {
	valueOfCmd := reflect.ValueOf(cmd)

	for valueOfCmd.Kind() == reflect.Interface || valueOfCmd.Kind() == reflect.Pointer {
		valueOfCmd = valueOfCmd.Elem()
	}

	if valueOfCmd.Kind() == reflect.Struct {
		return valueOfCmd, nil
	}

	return reflect.Value{}, fmt.Errorf("Command does not contain a valid struct")
}

func getNameOfCommand(cmd Command) (string, error) {
	valueOfCmd, err := inspectAndAccessValueOfCommand(cmd)
	if err != nil {
		return "", err
	}
	nameOfCmd := valueOfCmd.Type().Name()
	nameOfCmd = strings.ToLower(nameOfCmd)
	return nameOfCmd, nil
}
