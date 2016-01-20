package main

import "io/ioutil"
import "fmt"

func main() {
	filedata, err := ioutil.ReadFile("dna.txt")
	
	var dnastring string
	if err == nil {
		dnastring = string(filedata)
	}

	var pair string = ""
	for _, char := range dnastring {
		switch char {
			case 'A':
				pair = "T" + pair
			case 'T':
				pair = "A" + pair
			case 'C':
				pair = "G" + pair
			case 'G':
				pair = "C" + pair
		}
	}
	fmt.Println(pair)
}
