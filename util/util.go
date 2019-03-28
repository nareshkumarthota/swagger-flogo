package util

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"text/template"
)

// ExecuteTemplate creates required file from given inputs
func ExecuteTemplate(convType string, data interface{}) (bool, error) {

	var t *template.Template

	switch convType {
	case "flogoapiapp":
		t = template.Must(template.New("top").Parse(flogoAPITemplate))
	default:
		t = template.Must(template.New("top").Parse(flogoAPITemplate))
	}

	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		log.Fatal("error while rendering template from data: ", err)
	}
	s := buf.String()

	createFileWithContent("flogoapiapp.go", s)

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