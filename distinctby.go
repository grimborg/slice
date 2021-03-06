package slice

import "github.com/clipperhouse/typewriter"

var distinctBy = &typewriter.Template{
	Name: "DistinctBy",
	Text: `
// DistinctBy returns a new {{.SliceName}} whose elements are unique, where equality is defined by a passed func. See: http://clipperhouse.github.io/gen/#DistinctBy
func (rcv {{.SliceName}}) DistinctBy(equal func({{.Type}}, {{.Type}}) bool) (result {{.SliceName}}) {
	for _, v := range rcv {
		eq := func(_app {{.Type}}) bool {
			return equal(v, _app)
		}
		if !result.Any(eq) {
			result = append(result, v)
		}
	}
	return result
}
`}
