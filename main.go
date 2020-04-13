package main

import (
	"fmt"
	"os"
	"proj-gen/models"
)

var workspace string

func getPathFromWorkspace(){
	ws, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	workspace = ws
}

func CreateProject(project *models.Project){
	
}

func main(){
	getPathFromWorkspace()
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Path to descriptionfile not given")
		return
	}
	path := args[:1]
	var project models.Project
	project.Workspace = workspace
	xmlFile, err := os.Open(path[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer xmlFile.Close()
	if err = project.ParseFile(xmlFile); err != nil {
		fmt.Println(err)
	}
	project.Print()
}
