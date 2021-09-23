package main

import (
	"errors"
	"fmt"
	"strings"
)

type SCommands struct {
	APIFunctions map[string]SFunctions
}

func (Command *SCommands) Help(arguments ...[]string) error {
	help := "commands: [help, list, find, search, doc]"
	
	fmt.Println(help)
	
	return nil
}

func (Command *SCommands) List(arguments ...[]string) error {
	for fn := range Command.APIFunctions {
		fmt.Println(fn)
	}
	
	return nil
}

func (Command *SCommands) Find(arguments ...[]string) error {
	if len(arguments) <= 0 {
		return errors.New("not enough arguments supplied")
	}
	
	for fn := range Command.APIFunctions {
		if len(fn) >= len(arguments[0][0]) {
			if strings.Contains(strings.ToLower(fn)[0:len(arguments[0][0])], arguments[0][0]) {
				fmt.Println(fn)
			}
		}
	}
	
	return nil
}

func (Command *SCommands) DOC(arguments ...[]string) error {
	if len(arguments) <= 0 {
		return errors.New("not enough arguments supplied")
	}
	
	for fn := range Command.APIFunctions {
		if arguments[0][0] == strings.ToLower(fn) {
			arguments[0][0] = fn
		}
	}
	
	for fn := range Command.APIFunctions {
		if len(fn) >= len(arguments[0][0]) {
			function := Command.APIFunctions[fn]
			if strings.Contains(strings.ToLower(fn)[0:len(arguments[0][0])], arguments[0][0]) {
				if (len(function.Input) > 0 || len(function.Output) > 0) {
					fmt.Println(fn)
					
					for j := range function.Input {
						fmt.Printf("[%d] %s (%s) - %s\n", j, function.Input[j].Name, function.Input[j].Type, function.Input[j].Desc)
					}
					
					for j := range function.Output {
						fmt.Printf("Return %s (%s) - %s\n", function.Output[j].Name, function.Output[j].Type, function.Output[j].Desc)
					}
					
					fmt.Println()
				} else {
					fmt.Println("no input/output")
				}
			}
		}
	}
	
	return nil
}