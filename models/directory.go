package models

import (
	"encoding/xml"
	"fmt"
)

type Directory struct {
	XMLName xml.Name `xml:"dir"`
	Name string `xml:"name,attr"`
	Files []File `xml:"file"`
	Dirs []Directory `xml:"dir"`
	Prefix string
}

func (dir *Directory) Create() error{
	return nil
}

func (dir *Directory) CreateFiles() error{
	for _, file := range dir.Files{
		file.Create(dir.Prefix+"/"+dir.Name)
	}
	return nil
}

func (dir *Directory) CreateDirectories() error {
	for _, dir := range dir.Dirs{
		dir.Create()
	}
	return nil
}

func (dir *Directory) Print(depth int) {
	for i:=0; i < depth; i ++ {
		fmt.Print("    ")
	}
	fmt.Println("/", dir.Name)

	for _, file := range dir.Files {
		file.Print(depth+1)
	}

	for _, dirs := range dir.Dirs {
		dirs.Print(depth+1)
	}
}