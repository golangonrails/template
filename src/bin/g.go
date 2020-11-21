package main

import (
	"app/bin/gen"
	"fmt"
	"log"
	"os"
)

func help() {
	fmt.Println(`
			Code Generator

Option:
	migration   <name>                     # Generate Migration File By Name
	seed        <name>                     # Generate Seed File By Name
`)
}

func at(arr []string, index int) string {
	if len(arr) > index {
		return os.Args[index]
	}
	return ""
}

func main() {
	if action := at(os.Args, 1); action != "" {
		if f := (map[string]func(){
			"migration": func() {
				if fileName := at(os.Args, 2); fileName != "" {
					gen.Migration(fileName)
				} else {
					log.Printf("Command `%v`: need name", action)
				}
			},
			"seed": func() {
				if fileName := at(os.Args, 2); fileName != "" {
					gen.Seed(fileName)
				} else {
					log.Printf("Command `%v`: need name", action)
				}
			},
		})[action]; f != nil {
			f()
			return
		} else {
			log.Printf("Command `%v` not found", action)
			return
		}
	}
	help()
}
