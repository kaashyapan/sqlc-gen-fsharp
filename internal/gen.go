package fsharp

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"strings"

	plugin "github.com/tabbed/sqlc-go/codegen"

	"github.com/kaashyapan/sqlc-gen-fsharp/internal/core"
	"github.com/kaashyapan/sqlc-gen-fsharp/internal/tmpl"
)

func Generate(ctx context.Context, req *plugin.Request) (*plugin.Response, error) {
	conf, err := core.MakeConfig(req)

	enums := core.BuildEnums(req)
	structs := core.BuildDataClasses(conf, req)
	queries, err := core.BuildQueries(req, structs)
	if err != nil {
		return nil, err
	}

	i := &core.Importer{
		Settings:    req.Settings,
		Enums:       enums,
		DataClasses: structs,
		Queries:     queries,
	}

	core.DefaultImporter = i

	tctx := core.TmplCtx{
		Settings:      req.Settings,
		Q:             `"""`,
		Package:       conf.Package,
		Queries:       queries,
		Enums:         enums,
		DataClasses:   structs,
		SqlcVersion:   req.SqlcVersion,
		Configuration: conf,
	}

	output := map[string]string{}

	execute := func(name string, f func(io.Writer, core.TmplCtx) error) error {
		var b bytes.Buffer
		w := bufio.NewWriter(&b)
		tctx.SourceName = name
		err := f(w, tctx)
		w.Flush()
		if err != nil {
			return err
		}
		if !strings.HasSuffix(name, ".fs") {
			name += ".fs"
		}
		output[name] = core.Format(b.String())
		return nil
	}

	if err := execute("Models.fs", tmpl.Models); err != nil {
		return nil, err
	}
	if err := execute("Readers.fs", tmpl.Reader); err != nil {
		return nil, err
	}
	if err := execute("Queries.fs", tmpl.SQL); err != nil {
		return nil, err
	}

	resp := plugin.CodeGenResponse{}

	for filename, code := range output {
		resp.Files = append(resp.Files, &plugin.File{
			Name:     filename,
			Contents: []byte(code),
		})
	}

	return &resp, nil
}
