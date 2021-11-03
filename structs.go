package main

import "encoding/xml"

type SInput struct {
	Name     string `xml:"name,attr"`
	Type     string `xml:"type,attr"`
	Optional string `xml:"optional,attr"`
	Desc     string `xml:"desc,attr"`
}

type SOutput struct {
	Name string `xml:"name,attr"`
	Type string `xml:"type,attr"`
	Desc string `xml:"desc,attr"`
}

type SFunctions struct {
	Name   string    `xml:"name,attr"`
	Output []SOutput `xml:"output"`
	Input  []SInput  `xml:"input"`
}

type SAPI struct {
	XMLName  xml.Name     `xml:"api"`
	Function []SFunctions `xml:"function"`
}