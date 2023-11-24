package templates

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/flosch/pongo2/v6"
)

type Templates map[string]*pongo2.Template

// for now we dont error
func (views Templates) Add(name string, template *pongo2.Template) {

	if name == "" {
		return
	}

	if template == nil {
		return
	}

	views[name] = template
}

func (views Templates) Render(name string, w io.Writer, c pongo2.Context) error {

	if _, ok := views[name]; !ok {
		return fmt.Errorf("template %s not found", name)
	}

	return views[name].ExecuteWriter(c, w)
}

func (views Templates) Dbg() {

	for k := range views {
		fmt.Printf("template=%s\n", k)
	}
}

func LoadTemplates(dirpath string) *Templates {
	templs := make(Templates)

	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		fmt.Printf("path=%v\n", path)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if info.IsDir() {
			return nil
		}

		parent := filepath.Base(filepath.Dir(path))

		if parent == dirpath || parent == "layout" {
			fmt.Println("is dirpath")
			return nil
		}

		filename := filepath.Base(path)

		name := filename[:len(filename)-len(filepath.Ext(filename))]

		tname := parent + "_" + name

		templ := pongo2.Must(pongo2.FromFile(path))

		fmt.Printf("add tname=%v\n", tname)
		templs.Add(tname, templ)

		return nil
	})

	if err != nil {
		panic(err)
	}

	return &templs
}

