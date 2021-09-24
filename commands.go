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
	help := `commands: [help, list, find, search, doc]
help usage: help (command)
	`
	if len(arguments) >= 1 {
		commands := make(map[string] string);
		commands["list"] = "lists all functions"
		commands["find"] = "find a function via full function name or a partial name"
		commands["search"] = "alias of find"
		commands["doc"] = `returns the documentation for the function
example 1: doc register
	found: RegisterTool
	[0] id (string) - Tool unique identifier
	[1] name (string) - Tool name to show in hud
	[2] file (string) - Path to vox file
	
example 2: doc play
	found: PlayMusic
	[0] path (string) - Music path

	found: PlaySound
	[0] handle (number) - Sound handle
	[1] pos (table) - World position as vector. Default is player position.
	[2] volume (number) - Playback volume. Default is 1.0

	found: PlayLoop
	[0] handle (number) - Loop handle
	[1] pos (table) - World position as vector. Default is player position.
	[2] volume (number) - Playback volume. Default is 1.0
	
example 3: doc IsShapeVisible
	found: IsShapeVisible
	[0] handle (number) - Shape handle
	[1] maxDist (number) - Maximum visible distance
	[2] rejectTransparent (boolean) - See through transparent materials. Default false.
	Return visible (boolean) - Return true if shape is visible
		`
		
		_, exists := commands[arguments[0][0]]
		if exists {
			fmt.Println(commands[arguments[0][0]])
		} else {
			fmt.Println("please run 'help'")
		}
	} else {
		fmt.Println(help)
	}
	
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
			if strings.EqualFold(strings.ToLower(fn), strings.ToLower(arguments[0][0])) || strings.Contains(strings.ToLower(fn)[0:len(arguments[0][0])], arguments[0][0]) {
				fmt.Println("found:", fn)
				if (len(function.Input) > 0 || len(function.Output) > 0) {
					
					for j := range function.Input {
						fmt.Printf("[%d] %s (%s) - %s\n", j, function.Input[j].Name, function.Input[j].Type, function.Input[j].Desc)
					}
					
					for j := range function.Output {
						fmt.Printf("Return %s (%s) - %s\n", function.Output[j].Name, function.Output[j].Type, function.Output[j].Desc)
					}
					
					fmt.Println()
				} else {
					fmt.Println("no input/output")
					break
				}
			}
		}
	}
	
	return nil
}