package tmpl

import (
	"fmt"
	"io"

	"github.com/kaashyapan/sqlc-gen-fsharp/internal/core"
)

func KtModels(w io.Writer, dot core.KtTmplCtx) (err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			var ok bool
			if err, ok = recovered.(error); !ok {
				panic(recovered)
			}
		}
	}()
	return ktModelsTmpl(w, dot)
}

func ktModelsTmpl(w io.Writer, dot core.KtTmplCtx) error {
	_, _ = io.WriteString(w, "// Code generated by sqlc. DO NOT EDIT.\n// versions:\n//   sqlc ")
	_, _ = io.WriteString(w, dot.SqlcVersion)
	_, _ = io.WriteString(w, "\n\npackage ")
	_, _ = io.WriteString(w, dot.Package)
	_, _ = io.WriteString(w, "\n\n")
	if eval := core.Imports(dot.SourceName); len(eval) != 0 {
		for _, dot := range eval {
			_ = dot
			_, _ = io.WriteString(w, "\n")
			if eval := dot; len(eval) != 0 {
				for _, dot := range eval {
					_ = dot
					_, _ = io.WriteString(w, "import ")
					_, _ = io.WriteString(w, dot)
					_, _ = io.WriteString(w, "\n")
				}
			}
			_, _ = io.WriteString(w, "\n")
		}
	}
	_, _ = io.WriteString(w, "\n\n")
	if eval := dot.Enums; len(eval) != 0 {
		for _, dot := range eval {
			_ = dot
			_, _ = io.WriteString(w, "\n")
			if eval := dot.Comment; len(eval) != 0 {
				_, _ = io.WriteString(w, core.DoubleSlashComment(dot.Comment))
			}
			_, _ = io.WriteString(w, "\nenum class ")
			_, _ = io.WriteString(w, dot.Name)
			_, _ = io.WriteString(w, "(val value: String) {")
			if eval := dot.Constants; len(eval) != 0 {
				for _Vari, _Vare := range eval {
					_ = _Vari
					dot := _Vare
					_ = dot
					if eval := _Vari; eval != 0 {
						_, _ = io.WriteString(w, ",")
					}
					_, _ = io.WriteString(w, "\n  ")
					_, _ = io.WriteString(w, dot.Name)
					_, _ = io.WriteString(w, "(\"")
					_, _ = io.WriteString(w, dot.Value)
					_, _ = io.WriteString(w, "\")")
				}
			}
			_, _ = io.WriteString(w, ";\n\n  companion object {\n    private val map = ")
			_, _ = io.WriteString(w, dot.Name)
			_, _ = io.WriteString(w, ".values().associateBy(")
			_, _ = io.WriteString(w, dot.Name)
			_, _ = io.WriteString(w, "::value)\n    fun lookup(value: String) = map[value]\n  }\n}\n")
		}
	}
	_, _ = io.WriteString(w, "\n\n")
	if eval := dot.DataClasses; len(eval) != 0 {
		for _, dot := range eval {
			_ = dot
			_, _ = io.WriteString(w, "\n")
			if eval := dot.Comment; len(eval) != 0 {
				_, _ = io.WriteString(w, core.DoubleSlashComment(dot.Comment))
			}
			_, _ = io.WriteString(w, "\ndata class ")
			_, _ = io.WriteString(w, dot.Name)
			_, _ = io.WriteString(w, " (")
			if eval := dot.Fields; len(eval) != 0 {
				for _Vari, _Vare := range eval {
					_ = _Vari
					dot := _Vare
					_ = dot
					if eval := _Vari; eval != 0 {
						_, _ = io.WriteString(w, ",")
					}
					if eval := dot.Comment; len(eval) != 0 {
						_, _ = io.WriteString(w, "\n  ")
						_, _ = io.WriteString(w, core.DoubleSlashComment(dot.Comment))
					} else {
					}
					_, _ = io.WriteString(w, "\n  val ")
					_, _ = io.WriteString(w, dot.Name)
					_, _ = io.WriteString(w, ": ")
					_, _ = fmt.Fprint(w, dot.Type)
				}
			}
			_, _ = io.WriteString(w, "\n)\n")
		}
	}
	_, _ = io.WriteString(w, "\n\n")
	return nil
}
