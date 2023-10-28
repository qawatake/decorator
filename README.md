# decorator

Library `decorator` wraps analyzers to supplement the results with additional details.

Before

```
internal/example/example.go:11:16: nil dereference in field selection
```

After

```
internal/example/example.go:11:16: 😱 nil dereference in field selection (nilness)
```

## How to use

```go
package main

import (
  "github.com/qawatake/decorator"

  "golang.org/x/tools/go/analysis"
  "golang.org/x/tools/go/analysis/passes/nilness"
  "golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
  unitchecker.Main(
    decorator.With(
      func(a *analysis.Analyzer, d analysis.Diagnostic) analysis.Diagnostic {
        d.Message = "😱 " + d.Message + " (" + a.Name + ")"
        return d
      },
    )(nilness.Analyzer),
  )
}
```
