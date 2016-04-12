package xml2struct

import (
	"container/list"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
)

type Ele struct {
	Name  string
	child map[string]int //count the child number
}

func Parserxml(f string) (r map[string]map[string]int) {

	var res map[string]map[string]int = make(map[string]map[string]int)

	flag.Parse()

	xmlFile, err := os.Open(f)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	//A Token is an interface holding one of the token types: StartElement,
	// EndElement, CharData, Comment, ProcInst, or Directive.
	decoder := xml.NewDecoder(xmlFile)
	total := 0

	l := list.New()
	for {

		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()

		// Inspect the type of the token just read.
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			var e Ele

			e.Name = se.Name.Local
			e.child = make(map[string]int)

			if l.Len() > 0 {
				var lname string
				var curnum int
				if inst, ok := l.Back().Value.(Ele); ok {
					inst.child[e.Name]++
					lname = inst.Name
					curnum = inst.child[e.Name]
				}
				fmt.Printf("size:%d %s %s %d\n", l.Len(), lname, e.Name, curnum)
			}

			l.PushBack(e)

		case xml.EndElement:

			last := l.Back()

			if inst, ok := last.Value.(Ele); ok {
				_, ok := res[inst.Name]
				if !ok {

					res[inst.Name] = make(map[string]int)
				}

				for k, v := range inst.child {
					res[inst.Name][k] = v
				}

			}
			l.Remove(last)

		default:

		}
		total++

	}
	fmt.Printf("********************\n")

	return res
}
