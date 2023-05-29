package tmpl

import (
	"io"
	"strconv"

	"github.com/kaashyapan/sqlc-gen-fsharp/internal/core"
	"github.com/kaashyapan/sqlc-gen-fsharp/internal/templates"
)

func SQL(w io.Writer, dot core.TmplCtx) (err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			var ok bool
			if err, ok = recovered.(error); !ok {
				panic(recovered)
			}
		}
	}()

	if dot.Settings.Engine == "mysql" {
		return nil
	}

	templates.WriteQueries(w, dot)
	//return fsSQLTmpl(w, dot)
	return nil
}

func fsSQLTmpl(w io.Writer, dot core.TmplCtx) error {
	ctx := dot

	if eval := dot.Queries; len(eval) != 0 {
		for _, dot := range eval {
			_ = dot

			if dot.Ret.EmitStruct() {
				_, _ = io.WriteString(w, "data class ")
				_, _ = io.WriteString(w, dot.Ret.Type())
				_, _ = io.WriteString(w, " (")
				if eval := dot.Ret.Struct.Fields; len(eval) != 0 {
					for i, dot := range eval {
						_ = dot
						if i > 0 {
							_, _ = io.WriteString(w, ",")
						}
						_, _ = io.WriteString(w, "\n  val ")
						_, _ = io.WriteString(w, dot.Name)
						_, _ = io.WriteString(w, ": ")
						_, _ = io.WriteString(w, dot.Type.String())
					}
				}
				_, _ = io.WriteString(w, "\n)\n\n")
			}
		}
	}

	if dot.Settings.Engine == "mysql" {
		return nil
	}

	if eval := dot.Queries; len(eval) != 0 {
		_, _ = io.WriteString(w, "\ntype Queries(conn: string) = ")
		_, _ = io.WriteString(w, "\n    ")

		for _, dot := range eval {
			_ = dot
			if dot.Cmd == ":one" {
				if eval := dot.Comments; len(eval) != 0 {
					_, _ = io.WriteString(w, "\n")
					for _, dot := range eval {
						_ = dot
						_, _ = io.WriteString(w, "\n//")
						_, _ = io.WriteString(w, dot)
					}
				}

				_, _ = io.WriteString(w, "\n\n    member this.")
				_, _ = io.WriteString(w, dot.MethodName)
				_, _ = io.WriteString(w, "(")
				_, _ = io.WriteString(w, dot.Arg.Args())
				_, _ = io.WriteString(w, ") =")

				_, _ = io.WriteString(w, "\n      ")
				_, _ = io.WriteString(w, core.ExecCommand(ctx, dot))
				_, _ = io.WriteString(w, " ")
				if dot.Ret.IsStruct() {
					_, _ = io.WriteString(w, "\nEmit - "+strconv.FormatBool(dot.Ret.EmitStruct()))

				} else {
					_, _ = io.WriteString(w, "\n//")
					_, _ = io.WriteString(w, "\nName - "+dot.Ret.Name)
					_, _ = io.WriteString(w, "\nType name - "+dot.Ret.Typ.Name)
					_, _ = io.WriteString(w, "\nEmit - "+strconv.FormatBool(dot.Ret.Emit))
					_, _ = io.WriteString(w, "\nType datatype - "+dot.Ret.Typ.DataType)
					_, _ = io.WriteString(w, "\nDB Type name - "+dot.Ret.Typ.DbName)
					_, _ = io.WriteString(w, "\nType libtype - "+dot.Ret.Typ.LibTyp)
					_, _ = io.WriteString(w, "\nType readertype - "+dot.Ret.Typ.ReaderTyp)

				}

			}
			if dot.Cmd == ":many" {
				if eval := dot.Comments; len(eval) != 0 {
					_, _ = io.WriteString(w, "\n")
					for _, dot := range eval {
						_ = dot
						_, _ = io.WriteString(w, "\n//")
						_, _ = io.WriteString(w, dot)
					}
				}
				_, _ = io.WriteString(w, "\n\n    member this.")
				_, _ = io.WriteString(w, dot.MethodName)
				_, _ = io.WriteString(w, "(")
				_, _ = io.WriteString(w, dot.Arg.Args())
				_, _ = io.WriteString(w, ") =")

				_, _ = io.WriteString(w, "\n      ")
				_, _ = io.WriteString(w, core.ExecCommand(ctx, dot))
				_, _ = io.WriteString(w, " ")

			}
			if dot.Cmd == ":exec" {
				if eval := dot.Comments; len(eval) != 0 {
					_, _ = io.WriteString(w, "\n")
					for _, dot := range eval {
						_ = dot
						_, _ = io.WriteString(w, "\n//")
						_, _ = io.WriteString(w, dot)
					}
				}

				_, _ = io.WriteString(w, "\n\n    member this.")
				_, _ = io.WriteString(w, dot.MethodName)
				_, _ = io.WriteString(w, "(")
				_, _ = io.WriteString(w, dot.Arg.Args())
				_, _ = io.WriteString(w, ") =")

				conn := ctx.ConnString()
				for _, line := range conn {
					_, _ = io.WriteString(w, line)
				}

				pipelines := ctx.ConnPipeline(dot)
				for _, line := range pipelines {
					_, _ = io.WriteString(w, line)
				}

				_, _ = io.WriteString(w, "\n      ")
				_, _ = io.WriteString(w, core.ExecCommand(ctx, dot))

			}

		}
	}
	_, _ = io.WriteString(w, "\n\n")
	return nil
}
