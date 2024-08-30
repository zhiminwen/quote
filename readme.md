# Quote for quick string constructions

Construct a string slice as Ruby's quote
- quote.Word() ~ Ruby's %w
- quote.Heredoc() ~ Ruby's <<~ EOF
- quote.Template() ~ Ruby's embedded string #{var}
- quote.Line()

>Based on the suggestion, change the q to quote for a better package name

## Update
- switch to go mod
- use sprig v3

## Synposis

Get and import
```
go get github.com/zhiminwen/quote
```

import "github.com/zhiminwen/quote"

Sample:

```golang
var list []string

// slice of word
list = quote.Word(`master1 master2 master3`)

// slice from lines
list = quote.Line(`
  // Lines are trimmed with Left and Right Spaces
  # Empty line and comment line (started with # or //) are ignored
  line 1
  line 2
`)

// a heredoc and docf
doc := quote.HereDoc(`
  A block of heredoc with proper indent. (similar as ruby's <<~ )
    Simply a wrapper of github.com/MakeNowJust/heredoc heredoc.Doc(text)
`)

docf := quote.HereDocf(`
  A block of heredoc with proper indent. (similar as ruby's <<~ )
    Simply a wrapper of github.com/MakeNowJust/heredoc heredoc.Doc(text)
    Variable = %s
`, "Value")

// a shortcut for join commands
cmd := quote.Cmd(`
  // construct command lines. Same as Line(), join lines with " && " or other connector like ";"
  ls -ltr

  hostname
`, " && ")
//cmd == "ls -ltr && hostname"


//Construct a quick template based string using map[string]string
text := quote.Template(`This is a test for {{ .Name}} at {{ .Address }}`, map[string]string{
  Name: "test",
  Address: "Street 1",
})

```
