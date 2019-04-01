package util

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

// ExecuteTemplate creates required file from given inputs
func ExecuteTemplate(convType string, data interface{}) (bool, error) {

	var t *template.Template

	switch convType {
	case "flogoapiapp":
		t = template.Must(template.New("top").Parse(flogoAPITemplate))
	case "flogodescriptor":
		t = template.Must(template.New("top").Parse(flogoAppDescriptor))
	default:
		t = template.Must(template.New("top").Parse(flogoAPITemplate))
	}

	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		log.Fatal("error while rendering template from data: ", err)
	}
	s := buf.String()

	fileName := "flogoapiapp.go"
	switch convType {
	case "flogodescriptor":
		fileName = "flogodescriptor.json"
	}

	createFileWithContent(fileName, s)

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
	return strings.ReplaceAll(strings.ReplaceAll(path, "{", ":"), "}", "")
}
