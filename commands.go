package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type SCommands struct {
	APIFunctions map[string]SFunctions
}

const (
	columnWidth = 30
)

func (Command *SCommands) Help(arguments ...[]string) error {
	help := `Commands: [help, version, list, find, search, doc]
Help usage: help (command)
	`
	if len(arguments) >= 1 {
		commands := make(map[string] string);
		commands["version"] = "Displays current API Version"
		commands["list"] = "Lists all functions"
		commands["find"] = "Find a function via full function name or a partial name"
		commands["search"] = "Alias of find"
		commands["download"] = "Download the API for Offline Use"
		commands["doc"] = `Returns the documentation for the function
Example 1: doc register
[RegisterTool]
	[0] id (string) - Tool unique identifier
	[1] name (string) - Tool name to show in hud
	[2] file (string) - Path to vox file
	
Example 2: doc play
[PlayMusic]
	[0] path (string) - Music path

[PlaySound]
	[0] handle (number) - Sound handle
	[1] pos (table) - World position as vector. Default is player position.
	[2] volume (number) - Playback volume. Default is 1.0

[PlayLoop]
	[0] handle (number) - Loop handle
	[1] pos (table) - World position as vector. Default is player position.
	[2] volume (number) - Playback volume. Default is 1.0
	
Example 3: doc IsShapeVisible
[IsShapeVisible]
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
	var functions[] string
	for fn := range Command.APIFunctions {
		functions = append(functions, fn)
	}

	sort.Strings(functions)
	for fn := range functions {
		fmt.Println(functions[fn])
	}
	
	return nil
}

func (Command *SCommands) Find(arguments ...[]string) error {
	if len(arguments) <= 0 {
		return errors.New("not enough arguments supplied")
	}
	
	count := 0
	
	for fn := range Command.APIFunctions {
		if len(fn) >= len(arguments[0][0]) { // check if word from input matches first word in func name
			if strings.Contains(strings.ToLower(fn)[0:len(arguments[0][0])], arguments[0][0]) {
				count++
				fmt.Println(fn)
			} else if count == 0 { // no word starting with input so match any function with input
				if strings.Contains(strings.ToLower(fn), arguments[0][0]) {
					fmt.Println(fn)
				}
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
				fmt.Printf("[%s]\n", fn)
				if (len(function.Input) > 0 || len(function.Output) > 0) {
					for j := range function.Input {
						fmt.Printf("    [%d] %s (%s) - %s\n", j, function.Input[j].Name, function.Input[j].Type, function.Input[j].Desc)
					}
					
					for j := range function.Output {
						fmt.Printf("    Return %s (%s) - %s\n", function.Output[j].Name, function.Output[j].Type, function.Output[j].Desc)
					}
				} else {
					fmt.Println("     This function does not return anything")
					break
				}
			}
		}
	}
	
	return nil
}

func (Command *SCommands) Version(arguments ...[]string) error {
	util := Utilities{}
	version, err := util.GetVersion(); if err != nil {
		return errors.New("Error: Unable to Fetch Version")
	}
	
	fmt.Printf("API Version: %s\n", version)
	
	return nil
}

func (Command* SCommands) Download(arguments ...[]string) error {
	util := Utilities{}
	err := util.DownloadLatestXML(); if err != nil {
		return errors.New("Error: Unable to Download API")
	}
	
	fmt.Printf("Successfully Downloaded API\n")
	
	return nil
}