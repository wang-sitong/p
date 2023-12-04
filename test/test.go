package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"os"
)

func main() {
	data := map[string]interface{}{
		"name":  "John Doe",
		"age":   30,
		"email": "john.doe@example.com",
	}
	r(data)
	os.Exit(1)
}

func r(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v)
		return
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")

	if err != nil {
		fmt.Println(err)
		return
	}

	printColored(out.String(), color.FgRed, color.BgYellow)

}

func printColored(text string, attributes ...color.Attribute) {
	c := color.New(attributes...)
	c.Println(text)
}
