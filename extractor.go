package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func isError(e error) bool {
	if e != nil {
		panic(e)
	}
	return true
}

func main() {
	//Validate command line args
	if len(os.Args) < 2 {
		panic("Usage: extractor <model.json filepath>")
	}

	//UI component will write model.json file
	//Get filepath from command line args
	fmt.Println("Reading " + os.Args[1])
	modelJson, err := ioutil.ReadFile(os.Args[1])

	if !isError(err) {
		//Send to SchemaBuilder as HTTP POST request
		b := bytes.NewBufferString(modelJson)

		fmt.Println("Sending model to Schema Builder")
		http.Post("http://127.0.0.1:8081/", "text/json", b)
	}
}
