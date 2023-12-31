# decorator

[![Go Reference](https://pkg.go.dev/badge/github.com/qawatake/decorator.svg)](https://pkg.go.dev/github.com/qawatake/decorator)
[![test](https://github.com/qawatake/decorator/actions/workflows/test.yaml/badge.svg)](https://github.com/qawatake/decorator/actions/workflows/test.yaml)
[![codecov](https://codecov.io/gh/qawatake/decorator/graph/badge.svg?token=5BB1k2a601)](https://codecov.io/gh/qawatake/decorator)

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
