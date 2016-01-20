package main

import "io/ioutil"
import "fmt"
import "bytes"

func main() {
	filedata, err := ioutil.ReadFile("dna.txt")
	
	var dnastring string
	if err == nil {
		dnastring = string(filedata)
	}

	var buffer bytes.Buffer
	for _, char := range dnastring {
		if char == 'T' {
			buffer.WriteRune('U')
		} else {
			buffer.WriteRune(char)
		}
	}
	fmt.Println(&buffer)
}
