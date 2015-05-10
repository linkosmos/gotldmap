# gotldmap
Top Level Domain map

[![Build Status](https://travis-ci.org/linkosmosis/gotldmap.svg)](https://travis-ci.org/linkosmosis/gotldmap)
[![GoDoc](http://godoc.org/github.com/linkosmosis/gotldmap?status.svg)](http://godoc.org/github.com/linkosmosis/gotldmap)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](http://opensource.org/licenses/MIT)

### Functions
```go
	mark, err := gotldmap.FindByTld("com")
	tld, err := gotldmap.FindByMark(165)
```


### Usage
```go
package main

import (
	"fmt"

	"github.com/linkosmosis/gotldmap"
)

func main() {
	mark, err := gotldmap.FindByTld("com")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("TLD %s mark %d", "com", mark))
}
```
