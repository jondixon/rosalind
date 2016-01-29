package main

import "io/ioutil"
import "bytes"
import "fmt"

func main() {
	filedata, err := ioutil.ReadFile("data.txt")
	
	dnastring1 := ""
	dnastring2 := ""

	if err == nil {
		r := bytes.NewBuffer(filedata)
		dnastring1, err = r.ReadString('\n')
		dnastring2, err = r.ReadString('\n')
	}

	var hammdist int = 0
	for i, char := range dnastring1 {
		if dnastring1[i] != dnastring2[i] {
			hammdist++
		}
		fmt.Println(char)
	}
	fmt.Println(hammdist)
}
