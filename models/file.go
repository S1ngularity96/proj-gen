package models

import (
	"encoding/xml"
	"fmt"
)

type File struct{
	XMLName xml.Name `xml:"file"`
	Name string `xml:"name,attr"`
	Templatefile string `xml:"src,attr"`
}

func (file *File) Create(dir string) error{
	return nil
}

func (file *File) Print(depth int){
	for i:=0; i < depth; i ++ {
		fmt.Print("    ")
	}

	var link string = ""
	if file.Templatefile != "" {
		link = fmt.Sprintf("-> %s", file.Templatefile)
	}

	fmt.Println(file.Name, link)
}