package p

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func Pr(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v)
		return
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		fmt.Println(v)
		return
	}
	fmt.Println(out.String())
}

func Pd(v interface{}) {
	Pr(v)
	panic("")
}
