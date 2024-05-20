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
		boxes.CenterWidget{
			Child: boxes.ContainerWidget{
				Child: boxes.TextWidget{
					Msg: "Hello from a Widget",
				},
			},
		},
	)
}
```

Renders this UI:

```
                         ┌─────────────────────┐
                         │                     │
                         │ Hello from a Widget │
                         │                     │
                         └─────────────────────┘
```

## Dependencies


