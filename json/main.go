package main

import (
	"encoding/json"
	"fmt"

	"github.com/linkosmos/gotldmap"
)

func main() {
	binary, err := json.Marshal(gotldmap.Map)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(binary))
}
