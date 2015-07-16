package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

/*

File Type 1 is a config file which the key value paris have been moved from
the web.config file to a location outside of the web apps binaries.  The reason
for moving the configs is so the code can be updated without overwriting the
configs specific for the site

<appSettings>
  <add key="LogoUrl" value="~/img/logo_new.png"/>
  <add key="FavIcon" value="~/img/favicon_etss.ico"/>
</appSettings>
*/

/*
File Type 2 is the web.config format which exists in the
web application code.  Typically this config values are non specific
and are used for all web sites.

<configuration>
 <location allowOverride="true">
	<appSettings file="DAL.config">
		<add key="SiteName" value="Import/Export"/>
	</appSettings>
 </location>
</configuration>
*/

// Configuration is the top level element
type Configuration struct {
	XMLName  xml.Name   `xml:"configuration"`
	Location []Location `xml:"location"`
}

// Location is the secondary level element
type Location struct {
	XMLName     xml.Name      `xml:"location"`
	AppSettings []AppSettings `xml:"appSettings"`
}

// AppSettings contains a list of key value paris for .Net config files
// the key value pairs are used through out the application and read by
// .Net applications using .Net code
type AppSettings struct {
	XMLName xml.Name `xml:"appSettings"`
	Add     []Add    `xml:"add"`
}

//Add the annoying name for an App Setting KVP
type Add struct {
	XMLName xml.Name `xml:"add"`
	Key     string   `xml:"key,attr"`
	Value   string   `xml:"value,attr"`
}

func main() {

	var buf bytes.Buffer
	logger := log.New(&buf, "logger: ", log.Lshortfile)
	logger.Print("Hello, log file!")

	fmt.Print(&buf)
	var rootpath = "/Users/cn/Downloads/"

	if _, err := os.Stat(rootpath); os.IsNotExist(err) {
		fmt.Printf("%s does not exist\n", rootpath)
		checkError(err)
	}

	fileList := []string{}
	/* http://golang.org/pkg/path/filepath/
	filepath.Base - returns the name of the file
	filepath.Ext  - returns the filename extenstion

	*/
	err := filepath.Walk(rootpath, func(path string, f os.FileInfo, err error) error {
		var ext = filepath.Ext(path)

		if ext == ".config" {
			fileList = append(fileList, path)
		}
		return nil
	})
	checkError(err)

	// iterate over the list of files
	for _, file := range fileList {
		ReadXML(file, filepath.Base(file))
		fmt.Printf("\n")
	}
}

// ReadXML will read emit the key values in a .Net config file
func ReadXML(filepath string, filename string) {

	xmlFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()

	XMLdata, _ := ioutil.ReadAll(xmlFile)

	var c AppSettings
	var a Configuration

	fileType := 1
	if filename == "web.config" {
		fileType = 2
	}

	if fileType == 1 {
		xml.Unmarshal(XMLdata, &c)
		for _, k := range c.Add {
			fmt.Printf("Key: %s, Value: %s\n", k.Key, k.Value)
		}
	}

	if fileType == 2 {
		xml.Unmarshal(XMLdata, &a)
		for _, l := range a.Location {
			for _, as := range l.AppSettings {
				for _, k := range as.Add {
					fmt.Printf("Key: %s, Value: %s\n", k.Key, k.Value)
				}
			}
		}
	}
}

// checkError prints out the error to the console
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
