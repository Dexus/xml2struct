package xml2struct

import (
	"fmt"
	"testing"
)
import "bytes"

func TestParserxml(t *testing.T) {
	buf := bytes.NewBufferString("")
	buf.WriteString("1234")
	buf.WriteString("5678")
	fmt.Printf(buf.String())
}
