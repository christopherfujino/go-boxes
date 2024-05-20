# go-boxes

Box-y terminal UI framework.

## Examples

The following code:

```go
func main() {
  boxes.Run(
    boxes.Row{
      Children: boxes.Map(
        []string{
          "Hello, world!",
          "Hmm...",
          "Goodbye, world!",
        },
        func(s string) boxes.Widget {
          return boxes.Container{Child: boxes.Text{Msg: s}}
        },
      ),
    },
  )
}
```

Renders this UI:

```
┌───────────────┐┌────────┐┌─────────────────┐
│               ││        ││                 │
│ Hello, world! ││ Hmm... ││ Goodbye, world! │
│               ││        ││                 │
└───────────────┘└────────┘└─────────────────┘
```
