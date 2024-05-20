# go-boxes

Box model based terminal UI framework.

## Examples

The following code:

```go
package main

import (
  "github.com/christopherfujino/go-boxes"
)

func main() {
  boxes.Run(
    boxes.Center{
      Child: boxes.Container{
        Child: boxes.Text{
          Msg: "Hello, world!",
        },
      },
    },
  )
}
```

Renders this UI:

```
                         ┌───────────────┐
                         │               │
                         │ Hello, world! │
                         │               │
                         └───────────────┘
```

## Dependencies


