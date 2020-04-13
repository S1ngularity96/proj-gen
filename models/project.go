package models

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

var printOnly = func(r rune) rune {
	if unicode.IsPrint(r) {
		return r
	}
	return -1
}



type Project struct{
	XMLName xml.Name `xml:"project"`
	Projectname string `xml:"name,attr"`
	Description string `xml:"description,attr"`
	Files []File `xml:"file"`
	Dirs []Directory `xml:"dir"`
	Workspace string `xml:"-"`
}

func (project *Project) ParseFile(xmlFile *os.File) error{
	byteValue, err := ioutil.ReadAll(xmlFile)
	xmlData := []byte(strings.Map(printOnly, string(byteValue)))
	if err != nil {
		return err
	}
	err = xml.Unmarshal(xmlData, &project)
	if err != nil {
		return err
	}
	return nil
}

func (project *Project) Print(){
	fmt.Println("Name:", project.Projectname)
	fmt.Println("Description:", project.Description)

	for _, dirs := range project.Dirs{
		dirs.Print(0)
	}
	for _,files := range project.Files{
		files.Print(0)
	}
}