package main

import (
	"decorator"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/nilness"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(
		decorator.With(
			func(a *analysis.Analyzer, d analysis.Diagnostic) analysis.Diagnostic {
				// d.Message = d.Message + " (" + a.Name + ")"
				return d
			},
		)(nilness.Analyzer),
	)
}
