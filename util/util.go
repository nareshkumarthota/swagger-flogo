package util

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// ExecuteTemplate creates required file from given inputs
func ExecuteTemplate(conversionType, outFilePath string, data interface{}) (bool, error) {

	var t *template.Template

	funcMap := template.FuncMap{
		// The name "increment" is what the function will be called in the template text.
		"increment": func(i int) int {
			return i + 1
		},
	}

	var fileName string

	switch conversionType {
	case "flogoapiapp":
		fileName = "flogoapiapp.go"
		t = template.Must(template.New("top").Parse(flogoAPITemplate))
	case "flogodescriptor":
		fileName = "flogodescriptor.json"
		t = template.Must(template.New("top").Funcs(funcMap).Parse(flogoAppDescriptor))
	default:
		fileName = "flogoapiapp.go"
		t = template.Must(template.New("top").Parse(flogoAPITemplate))
	}

	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		log.Fatal("error while rendering template from data: ", err)
	}
	s := buf.String()

	// check for outfilepath if not exist create it
	if strings.Compare(outFilePath, ".") != 0 {
		_, err := os.Stat(outFilePath)
		if err != nil {
			fErr := os.MkdirAll(outFilePath, 0777)
			if fErr != nil {
				log.Fatal("unable to create out folder: ", outFilePath, fErr)
			}
		}
	}

	createFileWithContent(filepath.Join(outFilePath, fileName), s)

	// support File generation
	t = template.Must(template.New("top").Parse(supportFile))
	buf = &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		log.Fatal("error while rendering template from data: ", err)
	}
	s = buf.String()
	createFileWithContent(filepath.Join(outFilePath, "support.go"), s)

	return true, nil
}

func createFileWithContent(filename, content string) error {

	// Create a file on disk
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("error while creating file: %s", err.Error())
		return fmt.Errorf("error while creating file: %s", err.Error())
	}
	defer file.Close()

	// Open the file to write
	file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Printf("error while opening file: %s", err.Error())
		return fmt.Errorf("error while opening file: %s", err.Error())
	}

	// Write the Markdown doc to disk
	_, err = file.Write([]byte(content))
	if err != nil {
		log.Printf("error while writing Markdown to disk: %s", err.Error())
		return fmt.Errorf("error while writing Markdown to disk: %s", err.Error())
	}

	return nil
}

// ModifyPathSymbols modifies {} to : for flogo app usage
func ModifyPathSymbols(path string) string {
	return strings.Replace(strings.Replace(path, "{", ":", -1), "}", "", -1)
}
