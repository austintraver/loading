# Loading

<img height=200 width=526 src="https://user-images.githubusercontent.com/25112463/123889939-475bd380-d90b-11eb-9de4-d249dbf5c8f1.gif">

The `loading` package is a minimal package, built to provide command-line interfaces with the ability to present a
simple loading wheel to display to users on their console. As time passes, the console will be updated, simulating a
"spinning" animation.

## Installation

```shell
go get github.com/austintraver/loading
```

## Usage

```go
package main

import (
	"time"
	
	"github.com/austintraver/loading"
)

func main() {
    wheel := loading.New("%s Initializing...")
    wheel.Start()
    time.Sleep(5 * time.Second)
}
```

