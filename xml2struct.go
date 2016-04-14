package xml2struct

import (
	"bytes"
	"container/list"
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Ele struct {
	xml.Name
	child map[xml.Name]int //count the child number
	attr  map[xml.Name]int //
}

func GenerateStruct(res map[string]interface{}, prefix string, pt *os.File) string {
	buffer := bytes.NewBufferString("")

	for PNode, v := range res {
		inst1, ok := v.(Ele)
		if ok {
			fmt.Printf("cur hand " + PNode + "\n")
			line1 := fmt.Sprintf("type %s%s struct{\n", prefix, strings.Title(PNode))

			pt.WriteString(line1)
			line2 := fmt.Sprintf("\t%-20s\txml.Name `xml:\"%s %s\"`\n",
				"XMLName", inst1.Name.Space, PNode)
			pt.WriteString(line2)

			//generate child node string
			for S, v1 := range inst1.child {
				var line string
				if v1 == 1 {
					line = fmt.Sprintf("\t%-20s\t%s%s\t`xml:\"%s %s\"`\n",
						strings.Title(S.Local), prefix, strings.Title(S.Local), S.Space, S.Local)
				} else {
					line = fmt.Sprintf("\t%-20s\t[]%s%s\t`xml:\"%s %s\"`\n",
						strings.Title(S.Local), prefix, strings.Title(S.Local), S.Space, S.Local)
				}

				pt.WriteString(line)

			}
			//generate attr node string
			for S, _ := range inst1.attr {

				var line string

				line = fmt.Sprintf("\t%-20s\t%s%s\t`xml:\"%s %s,attr\"`\n",
					strings.Title(S.Local), prefix, strings.Title(S.Local), S.Space, S.Local)
				pt.WriteString(line)

			}

			pt.WriteString("}\n")
		}
	}
	return buffer.String()
}

func Parserxml(f string) (r map[string]interface{}) {

	var res map[string]interface{} = make(map[string]interface{})

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

			e.Name = se.Name
			e.child = make(map[xml.Name]int)
			e.attr = make(map[xml.Name]int)

			for i, k := range se.Attr {
				e.attr[k.Name] = i
			}

			if l.Len() > 0 {

				if inst, ok := l.Back().Value.(Ele); ok {
					inst.child[e.Name]++

				}
				//	fmt.Printf("size:%d %s %s %d\n", l.Len(), lname, e.Name, curnum)
			}

			l.PushBack(e)

		case xml.EndElement:

			last := l.Back()

			if inst, ok := last.Value.(Ele); ok {

				inst1, ok1 := res[inst.Name.Local].(Ele)

				if ok && ok1 { // existsing,need merge
					for i, k := range inst.child {
						if k == 2 {
							inst1.child[i] = k
						}

						//inst1.child[i]
					}

				} else { //exist merge into res
					res[inst.Name.Local] = last.Value
				}

				/*
					for k, v := range inst.child {
						if inst1, ok := last.Value.(Ele); ok {
							inst1.
						}
				}*/

			}
			l.Remove(last)

		default:

		}
		total++

	}
	fmt.Printf("********************\n")

	return res
}
