package decorator_test

import (
	"decorator"
	"testing"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/nilness"
)

// TestAnalyzer is a test for Analyzer.
func TestWith(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata,
		decorator.With(
			func(a *analysis.Analyzer, d analysis.Diagnostic) analysis.Diagnostic {
				d.Message = "decorated"
				return d
			},
		)(nilness.Analyzer),
		"a")
}
