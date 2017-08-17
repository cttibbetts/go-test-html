package main

import (
	"fmt"
	"log"
	"os"

	"io/ioutil"
	"path/filepath"

	rice "github.com/GeertJohan/go.rice"
	"github.com/ains/go-test-html/lib"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Incorrect command line arguments")
		fmt.Println("Usage: go-test-html [gotest_stdout_file] [gotest_stderr_file] [output_file]")
		os.Exit(1)
	}

	gotestStdoutFile := os.Args[1]
	gotestStderrFile := os.Args[2]
	outputDir := os.Args[3]

	gotestStdout, err := os.Open(gotestStdoutFile)
	check(err)

	gotestStderr, err := os.Open(gotestStderrFile)
	check(err)

	summary, err := lib.Parse(gotestStdout, gotestStderr)
	check(err)

	templateBox := rice.MustFindBox("template")
	html, err := lib.GenerateHTML(templateBox.MustString("template.html"), summary)
	check(err)

	js := templateBox.MustString("template.js")
	check(err)

	css := templateBox.MustString("template.css")
	check(err)

	err = os.Mkdir(outputDir, 1755)
	check(err)

	err = ioutil.WriteFile(outputDir+"/gotesthtml.html", []byte(html), 0644)
	check(err)

	err = ioutil.WriteFile(outputDir+"/gotesthtml.js", []byte(js), 0644)
	check(err)

	err = ioutil.WriteFile(outputDir+"/gotesthtml.css", []byte(css), 0644)
	check(err)

	outputFilePath, err := filepath.Abs(outputDir)
	check(err)

	fmt.Printf("Test results written to '%s'\n", outputFilePath)
}
