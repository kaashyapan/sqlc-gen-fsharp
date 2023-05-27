package core

import (
	"github.com/tabbed/sqlc-go/sdk"
)

func DoubleSlashComment(f string) string {
	return sdk.DoubleSlashComment(f)
}

func LowerTitle(f string) string {
	return sdk.LowerTitle(f)
}

var DefaultImporter *Importer

func Imports(filename string, pkgName string) []string {
	if DefaultImporter == nil {
		return nil
	}
	return DefaultImporter.Imports(filename, pkgName)
}
