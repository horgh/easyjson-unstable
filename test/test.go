package test

//go:generate easyjson test.go

import (
	"bytes"
	"encoding/json"
	"fmt"
)

//easyjson:json
type Foo struct {
	M map[string]string
}

func Test() {
	f := Foo{
		M: map[string]string{
			"a": "b",
			"c": "d",
		},
	}

	var buf2 []byte
	for i := 0; i < 1000; i++ {
		buf, err := json.Marshal(f)
		if err != nil {
			panic(err)
		}

		if buf2 != nil && !bytes.Equal(buf, buf2) {
			fmt.Printf(
				"unstable after %d iterations: %s vs. %s\n",
				i+1,
				string(buf2),
				string(buf),
			)
			return
		}

		buf2 = buf
	}
}
