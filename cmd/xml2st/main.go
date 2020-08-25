package main

import "fmt"
import "github.com/wxd237/xml2struct"
import "os"

func main() {
	fmt.Printf("Start\n")
	res := xml2struct.Parserxml("document.xml")
	restring := xml2struct.GenerateStruct(res, "MyStruct", os.Stdout)
	fmt.Print(restring)
}
