package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type Utilities struct{}

func (util *Utilities) GetXML() (*os.File, error) {
	/* Get current path */
	currentPath, err := os.Getwd(); if err != nil {
		return nil, fmt.Errorf("failed getting current path: %s", err)
	}

	/* Find File */
	var files[]string
	filepath.Walk(currentPath, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(".xml", f.Name()); if err == nil && r {
				files = append(files, f.Name())
				log.Println("file found:", f.Name())
			}
		}

		return nil
	})
	
	if len(files) <= 0 {
		return nil, errors.New("d")
	}

	/* Open File */
	file, err := os.Open(files[0]); if err != nil {
		return nil, fmt.Errorf("failed opening file \"%s\": %s", files[0], err)
	};
	
	return file, nil
}