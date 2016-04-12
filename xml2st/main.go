package main

import "fmt"
import "github.com/wxd237/xml2struct"
import "os"

func main() {
	fmt.Printf("asfsafa")
	res := xml2struct.Parserxml("../document.xml")
	restring := xml2struct.GenerateStruct(res, "Word", os.Stdout)
	fmt.Print(restring)
}
