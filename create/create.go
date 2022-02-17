package create

import (
	"log"
	"os"
	"path"
	"text/template"

	"github.com/fatih/color"
	"github.com/gosimple/slug"
)

type Article struct {
    Title string
}

var dir string

var filename string

const tpl = `# {{ .Title }}

`

func init() {
    log.SetPrefix("create: ")
    log.SetFlags(0)
}

func SetPath(p string) {
    _, err := os.Stat(p)
    if err != nil {
        log.Fatal(err)
    }
    if err == nil {
        dir = p
    }
}

func GetPath() string {
    return dir
}

func SetFilename(f string) {
    filename = f
}

func GetFilename() string {
    return filename
}

func GetUrl() string {
    return path.Join(GetPath(), slug.Make(GetFilename()) + ".md")
}

func Create() {
    if isFile() {
        return
    }

    template, err := template.New("article").Parse(tpl)
    if err != nil {
        log.Fatal(err)
    }

    f, err := os.Create(GetUrl())
    if err != nil {
        log.Fatal(err)
    }

    err = template.Execute(f, Article{GetFilename()})
    if err != nil {
        log.Fatal(err)
    }

    success := color.New(color.FgGreen)
    success.Println("Article was created successfully")
    success.Printf("File: %s\n", GetUrl())
}

func isFile() bool {
    if _, err := os.Stat(GetUrl()); err == nil {
        return true
    }

    return false
}
