package convergen

import (
	"go/ast"
	"strings"
)

// importNames is a map of a package path to a package name in a convergen setup file.
type importNames map[string]string

// newImportNames creates a new importNames instance.
func newImportNames(specs []*ast.ImportSpec) importNames {
	imports := make(importNames)
	for _, spec := range specs {
		pkgPath := strings.ReplaceAll(spec.Path.Value, `"`, "")
		imports[pkgPath] = spec.Name.Name
	}
	return imports
}

// lookupName looks up the map with the pkgPath and returns its corresponding name
// in the conversion setup file.
func (i importNames) lookupName(pkgPath string) (name string, ok bool) {
	name, ok = i[pkgPath]
	return
}