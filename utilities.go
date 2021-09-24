package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Utilities struct{}

var (
	version string = "0.0"
	apiHtml string = "https://www.teardowngame.com/modding/api.html"
	apiXml string = "https://www.teardowngame.com/modding/api.xml"
	file string
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
	
	fmt.Println("downloaded", file)
	
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
		err = util.DownloadLatestXML(); if err != nil {
			log.Println("using offline mode")
		}
	}

	if file == "api-.xml" {
		return nil, errors.New("unable to find teardown api xml, please make sure you are connected to the internet to download")
	}
	
	fmt.Println("file found:", file)
	
	/* Open File */
	file, err := os.Open(file); if err != nil {
		return nil, fmt.Errorf("failed opening file \"%s\": %s", file.Name(), err)
	};
	
	return file, nil
}