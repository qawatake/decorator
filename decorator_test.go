package decorator_test

import (
	"testing"

	"github.com/qawatake/decorator"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
	"golang.org/x/tools/go/analysis/passes/nilness"
)

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

func TestWithName(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata,
		decorator.WithName()(nilness.Analyzer),
		"b")
}
