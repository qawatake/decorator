package decorator

import (
	"golang.org/x/tools/go/analysis"
)

func With(decorator func(a *analysis.Analyzer, d analysis.Diagnostic) analysis.Diagnostic) func(a *analysis.Analyzer) *analysis.Analyzer {
	return func(a *analysis.Analyzer) *analysis.Analyzer {
		orgRun := a.Run
		a.Run = func(p *analysis.Pass) (interface{}, error) {
			orgReport := p.Report
			p.Report = func(d analysis.Diagnostic) {
				d = decorator(a, d)
				orgReport(d)
			}
			return orgRun(p)
		}
		return a
	}
}

func WithName() func(a *analysis.Analyzer) *analysis.Analyzer {
	return With(func(a *analysis.Analyzer, d analysis.Diagnostic) analysis.Diagnostic {
		d.Message = d.Message + " (" + a.Name + ")"
		return d
	})
}
