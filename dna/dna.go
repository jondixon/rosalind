package main

import "fmt"
import "io/ioutil"

func main () {
	filestuff, err := ioutil.ReadFile("dna.txt")
	
	var dnastring string
	if err == nil {
		dnastring = string(filestuff)
	}

	var a, c, t, g int = 0, 0, 0, 0
	for _, char := range dnastring {
		switch char {
			case 'A':
				a += 1
			case 'C':
				c += 1
			case 'T':
				t += 1
			case 'G':
				g += 1
		}
	}

	fmt.Println( a, " ", c, " ", g, " ", t, )
}
