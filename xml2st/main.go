package main

import "fmt"
import "github.com/wxd237/xml2struct"

func main() {
	fmt.Printf("asfsafa")
	res := xml2struct.Parserxml("../document.xml")
	for k, v := range res {
		for k1, v1 := range v {
			fmt.Printf("list:%s %s %d\n", k, k1, v1)

		}

	}
}
