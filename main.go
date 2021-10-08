package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	log "github.com/amoghe/distillog"
	"github.com/charmbracelet/lipgloss"
)

type SInput struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
	Optional string `xml:"optional,attr"`
	Desc string `xml:"desc,attr"`
}

type SOutput struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
	Desc string `xml:"desc,attr"`
}

type SFunctions struct {
	Name string `xml:"name,attr"`
	Output []SOutput `xml:"output"`
	Input []SInput `xml:"input"`
}

type SAPI struct {
	XMLName xml.Name `xml:"api"`
	Function []SFunctions `xml:"function"`
}

func main() {
	util := Utilities{}
	
	file, err := util.GetXML(); if err != nil {
		log.Errorln(err)
		fmt.Println("press any key to exit")
		bufio.NewReader(os.Stdin).ReadBytes('\n') 
		os.Exit(1)
	}; defer file.Close()

	/* Read api from file */
	byteValue, _ := ioutil.ReadAll(file)
	var api SAPI

	if err := xml.Unmarshal(byteValue, &api); err != nil {
		log.Errorln("unmarshal failed:", err.Error())
	}
	
	functions := make(map[string]SFunctions)
	
	for i := 0; i < len(api.Function); i++ {
		for i := 0; i < len(api.Function); i++ {
			functions[api.Function[i].Name] = api.Function[i]
		}
	}
	
	command := SCommands{}
	command.APIFunctions = functions
	commandMap := map[string] func(arguments ...[]string) error {
		"help": command.Help,
		"list": command.List,
		"find": command.Find,
		"search": command.Find,
		"doc": command.DOC,
		"version": command.Version,
	}
	
	HandleInput(commandMap)
}

var appStyle = lipgloss.NewStyle().Width(80)
var borderStyle = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
var specialStyle = lipgloss.AdaptiveColor{Light: "#43BF6D", Dark: "#73F59F"}

func HandleInput(commandMap map[string]func(arguments ...[]string) error) {
	scanner := bufio.NewScanner(os.Stdin)
	
	var input string
	
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