# gotldmap
Top Level Domain map

[![Build Status](https://travis-ci.org/linkosmos/gotldmap.svg)](https://travis-ci.org/linkosmos/gotldmap)
[![GoDoc](http://godoc.org/github.com/linkosmos/gotldmap?status.svg)](http://godoc.org/github.com/linkosmos/gotldmap)
[![BSD License](http://img.shields.io/badge/license-BSD-blue.svg)](http://opensource.org/licenses/BSD-3-Clause)

### Internet Assigned Numbers Authority

```
https://data.iana.org/TLD/tlds-alpha-by-domain.txt
```

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

	"github.com/linkosmos/gotldmap"
)

func main() {
	mark, err := gotldmap.FindByTld("com")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fmt.Sprintf("TLD %s mark %d", "com", mark))
}
```
