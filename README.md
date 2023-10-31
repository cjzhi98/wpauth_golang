# wpauth_golang

This is a Go package that provides functionality to check WordPress passwords.

## Installation

To install this package, you can use `go get`:

```
go get github.com/cjzhi98/wpauth_golang
```

## Usage

Here's a basic example of how to use this package:

```
package main

import (
    "fmt"
    "github.com/cjzhi98/wpauth_golang"
)

func main() {
    result := wpauth_golang.CheckPassword("password", "hash")
    fmt.Println(result)
}
```

In this example, replace "password" with the password you want to check and "hash" with the hashed password you want to compare it to.

##Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

License
MIT
