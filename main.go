package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	log "github.com/amoghe/distillog"
	"github.com/charmbracelet/lipgloss"
)

var (
	/* Styles */
	appStyle = lipgloss.NewStyle().Width(80)
	borderStyle = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	specialStyle = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}
	
	api SAPI
	command SCommands
	input string
	offline bool
)

func SetupAPIFile() error {
	util := Utilities{}
	
	file, err := util.GetXML(); if err != nil {
		return err
	}; defer file.Close()

	/* Read api from file */
	byteValue, _ := ioutil.ReadAll(file)

	if err := xml.Unmarshal(byteValue, &api); err != nil {
		log.Errorln("unmarshal failed:", err.Error())
	}
	
	functions := make(map[string]SFunctions)
	
	for i := 0; i < len(api.Function); i++ {
		for i := 0; i < len(api.Function); i++ {
			functions[api.Function[i].Name] = api.Function[i]
		}
	}
	
	command.APIFunctions = functions
	return nil
}

func SetupAPIOnline() error {
	r, err := http.Get("https://www.teardowngame.com/modding/api.xml"); if err != nil {
		return err
	}; defer r.Body.Close()

	err = xml.NewDecoder(r.Body).Decode(&api); if err != nil {
		fmt.Printf("failed decoding\n");
	}
	
	functions := make(map[string]SFunctions)
	
	for i := 0; i < len(api.Function); i++ {
		for i := 0; i < len(api.Function); i++ {
			functions[api.Function[i].Name] = api.Function[i]
		}
	}
	
	command.APIFunctions = functions
	return nil
}

func main() {
	err := SetupAPIOnline(); if err != nil {
		offline = true
		err := SetupAPIFile(); if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			fmt.Println("Press Enter to Exit")
			bufio.NewReader(os.Stdin).ReadBytes('\n') 
			os.Exit(1)	
		}
	}
	
	commandMap := map[string] func(arguments ...[]string) error {
		"help": command.Help,
		"list": command.List,
		"find": command.Find,
		"search": command.Find,
		"doc": command.DOC,
		"version": command.Version,
		"download": command.Download,
	}
	
	HandleInput(commandMap)
}

func HandleInput(commandMap map[string]func(arguments ...[]string) error) {
	scanner := bufio.NewScanner(os.Stdin)
	
	if input == "" {
		drawTitle()
		drawStatus()
		
		for scanner.Scan() {
			input = scanner.Text()
			input = strings.ToLower(input)
			
			arguments := strings.Fields(input)
			
			_, exists := commandMap[arguments[0]]
			if exists {
				if len(arguments) <= 1 {
					err := commandMap[arguments[0]](); if err != nil {
						fmt.Println(err.Error())
					}	
				} else {
					err := commandMap[arguments[0]](arguments[1:]); if err != nil {
						fmt.Println(err.Error())
					}
				}

			}

			drawStatus()
		}
	}
}