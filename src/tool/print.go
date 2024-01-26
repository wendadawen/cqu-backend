package tool

import "fmt"

func ShowContent(content interface{}) {
	sprintf := fmt.Sprintf("%v\n", content)
	fmt.Printf("%s", sprintf)
}
