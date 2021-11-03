package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Utilities struct{}

const (
	apiHtml string = "https://www.teardowngame.com/modding/api.html"
	apiXml string = "https://www.teardowngame.com/modding/api.xml"
)

var (
	file string
	version string
)

func (utilities *Utilities) GetVersion() (string, error) {
	r, err := http.Get(apiHtml); if err != nil {
		return "", err
	}
	
	doc, err := goquery.NewDocumentFromReader(r.Body); if err != nil {
		return "", err
	}
	
	title := doc.Find("h1").Text()
    version := regexp.MustCompile(`\((.*)\)`).FindStringSubmatch(title)[1]
	
	return version, nil
}

func (utilities *Utilities) DownloadLatestXML() error {	
	r, err := http.Get(apiXml); if err != nil {
		return err
	}; defer r.Body.Close()
	
	version, err = utilities.GetVersion(); if err != nil {
		return err
	}
	
	file = fmt.Sprintf("api-%s.xml", version)
	
	out, err := os.Create(file); if err != nil {
		return err
	}; defer out.Close()
	
	_, err = io.Copy(out, r.Body); if err != nil {
		return err
	}
	
	return nil
}

func (util *Utilities) GetXML() (*os.File, error) {
	version, _ := util.GetVersion()
	file = fmt.Sprintf("api-%s.xml", version)

	files, err := ioutil.ReadDir("."); if err != nil {
        return nil, err
    }
	
	found := false
    for _, f := range files {
		if strings.Contains(f.Name(), "api-") {
			found = true
			file = f.Name()
			break
		}
    }
	
	if !found {
		return nil, errors.New("Unable to find Teardown's API XML (api-*.xml), please make sure you are connected to the internet to download.")
	}
	
	/* Open File */
	file, err := os.Open(file); if err != nil {
		return nil, fmt.Errorf("Unable to Open \"%s\": %s", file.Name(), err)
	};

	return file, nil
}