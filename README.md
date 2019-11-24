# Zellers Congruence

A Go library which implements Zeller's Congruence, which determines the day of 
the week for an arbitrary date.

For a wonderful article about Zellers formula read this article:

- http://datagenetics.com/blog/november12019/index.html

# Installation

> $ go get github.com/umahmood/zellers

# Usage

```
package main

import (
    "fmt"

    "github.com/umahmood/zellers"
)

func main() {
    fmt.Println(zellers.Congruence(10, 15, 2019))
}
```
Output:
```
Tuesday
```

# Documentation

> http://godoc.org/github.com/umahmood/zellers

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).

