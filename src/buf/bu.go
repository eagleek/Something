package main

import (
	"bytes"
	"os"
	"fmt"
)

func main() {
	for i:=1;i<len(os.Args);i++{
		fmt.Printf("%s\n", intToString(os.Args[i]))
	}
}

func intToString(values string) string{

	var buf bytes.Buffer
	pre := len(values)%3	
	if pre ==0{
		pre =3
	}
	buf.WriteString(values[:pre])

	for i:= pre; i<len(values);i+=3{
		buf.WriteByte(',')
		buf.WriteString(values[i:i+3])
	}
	return buf.String()

}
 


