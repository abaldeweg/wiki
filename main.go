package main

import (
	"baldeweg/wiki/create"
	"flag"
	"fmt"
	"log"
	"os"
)

var dir string

var filename string

func init() {
    path, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }

    flag.StringVar(&dir, "path", path, "Specify the directory where the data should be stored.")
    flag.StringVar(&filename, "file", "", "Specify the filename.")
}

func main() {
    flag.Parse()

    create.SetPath(dir)

    if filename == "" {
        log.Fatal("No filename passed!")
    }
    create.SetFilename(filename)

    action := "new"
    if len(flag.Args()) >= 1 {
        action = flag.Args()[0]
    }

    switch action {
    case "new":
        create.Create()
    case "help":
        fmt.Println("baldeweg/wiki")
        fmt.Println("A baldeweg OpenSource project")
        fmt.Println("https://github.com/abaldeweg/wiki")
        fmt.Println("")
        fmt.Println("Commands")
        fmt.Println("mission [options] new - Adds a new wiki")
        fmt.Println("mission help - Shows the help")
    default:
        create.Create()
    }
}
