package core

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	easyjson "github.com/mailru/easyjson"
	plugin "github.com/tabbed/sqlc-go/codegen"
	"github.com/tabbed/sqlc-go/metadata"
	"github.com/tabbed/sqlc-go/sdk"

	"github.com/kaashyapan/sqlc-gen-fsharp/internal/inflection"
)

var fsIdentPattern = regexp.MustCompile("[^a-zA-Z0-9_]+")

type Constant struct {
	Name  string
	Type  string
	Value string
}

type Enum struct {
	Name      string
	Comment   string
	Constants []Constant
}

type Field struct {
	ID      int
	Name    string
	Type    fsType
	Comment string
}

type Struct struct {
	Table   plugin.Identifier
	Name    string
	Fields  []Field
	Comment string
}

type QueryValue struct {
	Emit   bool
	Name   string
	Struct *Struct
	Typ    fsType
}

func (v QueryValue) EmitStruct() bool {
	return v.Emit
}

func (v QueryValue) IsStruct() bool {
	return v.Struct != nil
}

func (v QueryValue) isEmpty() bool {
	return v.Typ == (fsType{}) && v.Name == "" && v.Struct == nil
}

func (v QueryValue) Type() string {
	if v.Typ != (fsType{}) {
		return v.Typ.String()
	}
	if v.Struct != nil {
		return v.Struct.Name
	}
	panic("no type for QueryValue: " + v.Name)
}

type Params struct {
	Struct  *Struct
	binding []int
}

func (v Params) isEmpty() bool {
	return len(v.Struct.Fields) == 0
}

func (v Params) Args() string {
	if v.isEmpty() {
		return ""
	}
	var requiredArgs []string
	var optionalArgs []string
	fields := v.Struct.Fields
	for _, f := range fields {

		if f.Type.IsNull {
			typ := strings.TrimSuffix(f.Type.String(), " option")
			optionalArgs = append(optionalArgs, "?"+f.Name+": "+typ)
		} else {
			requiredArgs = append(requiredArgs, f.Name+": "+f.Type.String())
		}

	}

	out := append(requiredArgs, optionalArgs...)
	return strings.Join(out, ", ")
}

func (v Params) Bindings(engine string) []string {
	var out []string

	for _, f := range v.Struct.Fields {
		item := fmt.Sprintf(`("%s", Sql.%s %s)`, f.Type.DbName, f.Type.LibTyp, f.Name)
		out = append(out, item)
	}

	return out
}

func indent(s string, n int, firstIndent int) string {
	lines := strings.Split(s, "\n")
	buf := bytes.NewBuffer(nil)
	for i, l := range lines {
		indent := n
		if i == 0 && firstIndent != -1 {
			indent = firstIndent
		}
		if i != 0 {
			buf.WriteRune('\n')
		}
		for i := 0; i < indent; i++ {
			buf.WriteRune(' ')
		}
		buf.WriteString(l)
	}
	return buf.String()
}

// A struct used to generate methods and fields on the Queries struct
type Query struct {
	ClassName    string
	Cmd          string
	Comments     []string
	MethodName   string
	FieldName    string
	ConstantName string
	SQL          string
	SourceName   string
	Ret          QueryValue
	Arg          Params
}

func fsEnumValueName(value string) string {
	id := strings.Replace(value, "-", "_", -1)
	id = strings.Replace(id, ":", "_", -1)
	id = strings.Replace(id, "/", "_", -1)
	id = fsIdentPattern.ReplaceAllString(id, "")
	return strings.ToUpper(id)
}

func BuildEnums(req *plugin.CodeGenRequest) []Enum {
	var enums []Enum
	for _, schema := range req.Catalog.Schemas {
		if schema.Name == "pg_catalog" || schema.Name == "information_schema" {
			continue
		}
		for _, enum := range schema.Enums {
			var enumName string
			if schema.Name == req.Catalog.DefaultSchema {
				enumName = enum.Name
			} else {
				enumName = schema.Name + "_" + enum.Name
			}
			e := Enum{
				Name:    dataClassName(enumName, req.Settings),
				Comment: enum.Comment,
			}
			for _, v := range enum.Vals {
				e.Constants = append(e.Constants, Constant{
					Name:  fsEnumValueName(v),
					Value: v,
					Type:  e.Name,
				})
			}
			enums = append(enums, e)
		}
	}
	if len(enums) > 0 {
		sort.Slice(enums, func(i, j int) bool { return enums[i].Name < enums[j].Name })
	}
	return enums
}

func dataClassName(name string, settings *plugin.Settings) string {
	if rename := settings.Rename[name]; rename != "" {
		return rename
	}
	out := ""
	for _, p := range strings.Split(name, "_") {
		out += sdk.Title(p)
	}
	return out
}

func memberName(name string, settings *plugin.Settings) string {
	return sdk.LowerTitle(dataClassName(name, settings))
}

func BuildDataClasses(conf Config, req *plugin.CodeGenRequest) []Struct {
	var structs []Struct
	for _, schema := range req.Catalog.Schemas {
		if schema.Name == "pg_catalog" || schema.Name == "information_schema" {
			continue
		}
		for _, table := range schema.Tables {
			var tableName string
			if schema.Name == req.Catalog.DefaultSchema {
				tableName = table.Rel.Name
			} else {
				tableName = schema.Name + "_" + table.Rel.Name
			}
			structName := dataClassName(tableName, req.Settings)
			if !conf.EmitExactTableNames {
				structName = inflection.Singular(inflection.SingularParams{
					Name:       structName,
					Exclusions: conf.InflectionExcludeTableNames,
				})
			}
			s := Struct{
				Table:   plugin.Identifier{Schema: schema.Name, Name: table.Rel.Name},
				Name:    structName,
				Comment: table.Comment,
			}
			for _, column := range table.Columns {
				s.Fields = append(s.Fields, Field{
					Name:    memberName(column.Name, req.Settings),
					Type:    makeType(req, column),
					Comment: column.Comment,
				})
			}
			structs = append(structs, s)
		}
	}
	return structs
}

type fsType struct {
	Name      string
	LibTyp    string
	ReaderTyp string
	DbName    string
	IsEnum    bool
	IsArray   bool
	IsNull    bool
	DataType  string
	Engine    string
}

func (t fsType) String() string {
	v := t.Name
	if t.IsArray {
		v = fmt.Sprintf("List<%s>", v)
	}
	return v
}

func (t fsType) IsTime() bool {
	return t.Name == "LocalDate" || t.Name == "LocalDateTime" || t.Name == "LocalTime" || t.Name == "OffsetDateTime"
}

func (t fsType) IsInstant() bool {
	return t.Name == "Instant"
}

func (t fsType) IsUUID() bool {
	return strings.ToLower(t.DataType) == "uuid"
}

func makeType(req *plugin.CodeGenRequest, col *plugin.Column) fsType {
	fstyp, readerTyp, libTyp, isEnum := fsInnerType(req, col)
	return fsType{
		Name:      fstyp,
		LibTyp:    libTyp,
		ReaderTyp: readerTyp,
		DbName:    col.Name,
		IsEnum:    isEnum,
		IsArray:   col.IsArray,
		IsNull:    !col.NotNull,
		DataType:  sdk.DataType(col.Type),
		Engine:    req.Settings.Engine,
	}
}

func fsInnerType(req *plugin.CodeGenRequest, col *plugin.Column) (string, string, string, bool) {
	// TODO: Extend the engine interface to handle types
	switch req.Settings.Engine {
	case "mysql":
		return mysqlType(req, col)
	case "postgresql":
		return postgresType(req, col)
	case "sqlite":
		return sqliteType(req, col)
	default:
		return "any", "any", "any", false
	}
}

type goColumn struct {
	id int
	*plugin.Column
}

func fsColumnsToStruct(req *plugin.CodeGenRequest, name string, columns []goColumn, namer func(*plugin.Column, int) string) *Struct {
	gs := Struct{
		Name: name,
	}
	idSeen := map[int]Field{}
	nameSeen := map[string]int{}
	for _, c := range columns {
		if _, ok := idSeen[c.id]; ok {
			continue
		}
		fieldName := memberName(namer(c.Column, c.id), req.Settings)
		if v := nameSeen[c.Name]; v > 0 {
			fieldName = fmt.Sprintf("%s_%d", fieldName, v+1)
		}
		field := Field{
			ID:   c.id,
			Name: fieldName,
			Type: makeType(req, c.Column),
		}
		gs.Fields = append(gs.Fields, field)
		nameSeen[c.Name]++
		idSeen[c.id] = field
	}
	return &gs
}

func fsArgName(name string) string {
	out := ""
	for i, p := range strings.Split(name, "_") {
		if i == 0 {
			out += strings.ToLower(p)
		} else {
			out += sdk.Title(p)
		}
	}
	return out
}

func fsParamName(c *plugin.Column, number int) string {
	if c.Name != "" {
		return fsArgName(c.Name)
	}
	return fmt.Sprintf("dollar_%d", number)
}

func fsColumnName(c *plugin.Column, pos int) string {
	if c.Name != "" {
		return c.Name
	}
	return fmt.Sprintf("column_%d", pos+1)
}

// HACK: jdbc doesn't support numbered parameters, so we need to transform them to question marks...
// But there's no access to the SQL parser here, so we just do a dumb regexp replace instead. This won't work if
// the literal strings contain matching values, but good enough for a prototype.
func jdbcSQL(s, engine string) (string, []string) {
	return s, nil
}

// Converts $1 to @id
func reformatSqlParamNames(q Query) string {
	rawQuery := q.SQL
	if q.Arg.isEmpty() {
		return rawQuery
	}

	if len(q.Arg.binding) > 0 {
		for _, idx := range q.Arg.binding {
			f := q.Arg.Struct.Fields[idx-1]
			token := `\$` + strconv.Itoa(idx+1) + `\b`
			regx := regexp.MustCompile(token)
			newToken := "@" + f.Type.DbName
			rawQuery = regx.ReplaceAllString(rawQuery, newToken)
		}
	} else {
		for i, f := range q.Arg.Struct.Fields {
			token := `\$` + strconv.Itoa(i+1) + `\b`
			regx := regexp.MustCompile(token)
			newToken := "@" + f.Type.DbName
			rawQuery = regx.ReplaceAllString(rawQuery, newToken)
		}
	}

	return rawQuery

}

// provide initial connection string
func (t TmplCtx) ConnString() []string {
	if t.Settings.Engine == "postgresql" {
		out := []string{"// https://www.connectionstrings.com/npgsql"}
		return out
	}
	if t.Settings.Engine == "sqlite" {
		out := []string{"// https://www.connectionstrings.com/sqlite-net-provider"}
		return out
	}
	return nil
}

// provide initial connection string
func (t TmplCtx) ConnPipeline(q Query) []string {
	out := []string{}
	argCnt := len(q.Arg.Bindings(t.Settings.Engine))

	if argCnt > 0 {
		paramstr := fmt.Sprintf("let parameters = [ %s ]", strings.Join(q.Arg.Bindings(t.Settings.Engine), "; "))
		out = append(out, paramstr)
		out = append(out, "")
	}

	out = append(out, "conn")
	out = append(out, "|> Sql.connect")
	out = append(out, "|> Sql.query Sqls."+q.ConstantName)

	if argCnt > 0 {
		out = append(out, "|> Sql.parameters parameters")
	}
	out = append(out, "|> Sql."+ExecCommand(t, q))

	return out
}

func parseInts(s []string) ([]int, error) {
	if len(s) == 0 {
		return nil, nil
	}
	var refs []int
	for _, v := range s {
		i, err := strconv.Atoi(strings.TrimPrefix(v, "$"))
		if err != nil {
			return nil, err
		}
		refs = append(refs, i)
	}
	return refs, nil
}

func BuildQueries(req *plugin.CodeGenRequest, structs []Struct) ([]Query, error) {
	qs := make([]Query, 0, len(req.Queries))
	for _, query := range req.Queries {
		if query.Name == "" {
			continue
		}
		if query.Cmd == "" {
			continue
		}
		if query.Cmd == metadata.CmdCopyFrom {
			return nil, errors.New("Support for CopyFrom in fsharp is not implemented")
		}

		ql, args := jdbcSQL(query.Text, req.Settings.Engine)
		refs, err := parseInts(args)
		if err != nil {
			return nil, fmt.Errorf("Invalid parameter reference: %w", err)
		}
		gq := Query{
			Cmd:          query.Cmd,
			ClassName:    sdk.Title(query.Name),
			ConstantName: sdk.LowerTitle(query.Name),
			FieldName:    sdk.Title(query.Name),
			MethodName:   sdk.LowerTitle(query.Name),
			SourceName:   query.Filename,
			SQL:          ql,
			Comments:     query.GetComments(),
		}

		var cols []goColumn
		for _, p := range query.Params {
			cols = append(cols, goColumn{
				id:     int(p.Number),
				Column: p.Column,
			})
		}
		params := fsColumnsToStruct(req, gq.ClassName+"Bindings", cols, fsParamName)
		gq.Arg = Params{
			Struct:  params,
			binding: refs,
		}

		if len(query.Columns) == 1 {
			c := query.Columns[0]
			gq.Ret = QueryValue{
				Name: "results",
				Typ:  makeType(req, c),
			}
		} else if len(query.Columns) > 1 {
			var gs *Struct
			var emit bool

			for _, s := range structs {
				if len(s.Fields) != len(query.Columns) {
					continue
				}
				same := true
				for i, f := range s.Fields {
					c := query.Columns[i]
					sameName := f.Name == memberName(fsColumnName(c, i), req.Settings)
					sameType := f.Type == makeType(req, c)
					sameTable := sdk.SameTableName(c.Table, &s.Table, req.Catalog.DefaultSchema)

					if !sameName || !sameType || !sameTable {
						same = false
					}
				}
				if same {
					gs = &s
					break
				}
			}

			if gs == nil {
				var columns []goColumn
				for i, c := range query.Columns {
					columns = append(columns, goColumn{
						id:     i,
						Column: c,
					})
				}
				gs = fsColumnsToStruct(req, gq.ClassName+"Row", columns, fsColumnName)
				emit = true
			}
			gq.Ret = QueryValue{
				Emit:   emit,
				Name:   "results",
				Struct: gs,
			}
		}

		gq.SQL = reformatSqlParamNames(gq)

		qs = append(qs, gq)
	}
	//sort.Slice(qs, func(i, j int) bool { return qs[i].MethodName < qs[j].MethodName })
	return qs, nil
}

type TmplCtx struct {
	Q           string
	Package     string
	Enums       []Enum
	DataClasses []Struct
	Queries     []Query
	Settings    *plugin.Settings
	SqlcVersion string
	// TODO: Race conditions
	SourceName string

	Configuration       Config
	EmitJSONTags        bool
	EmitPreparedQueries bool
	EmitInterface       bool
}

func makeReaderString(lookup map[string]Query, t TmplCtx) []string {
	var readers []string
	var fields []string

	for _, v := range lookup {
		fields = []string{}
		//cnt := len(v.Ret.Struct.Fields)
		for _, item := range v.Ret.Struct.Fields {
			field := fmt.Sprintf(`%s = r.%s "%s"`, sdk.Title(item.Name), item.Type.ReaderTyp, item.Type.DbName)
			fields = append(fields, field)
		}
		recstr := strings.Join(fields, " ; ")
		str := fmt.Sprintf(`let %sReader (r: RowReader) : %s = { %s.%s }`, sdk.LowerTitle(v.Ret.Type()), v.Ret.Type(), v.Ret.Type(), recstr)
		readers = append(readers, str)
	}
	return readers

}

func (v TmplCtx) ReaderSet() []string {

	lookup := map[string]Query{}
	if eval := v.Queries; len(eval) != 0 {
		for _, dot := range eval {
			_ = dot
			if dot.Cmd == ":one" || dot.Cmd == ":many" {
				if dot.Ret.IsStruct() {
					lookup[dot.Ret.Type()] = dot
				}
			}
		}
	}

	return makeReaderString(lookup, v)
}

func (v TmplCtx) ExtraModels() []string {

	lookup := map[string]Query{}

	if eval := v.DataClasses; len(eval) != 0 {
	}
	if eval := v.Queries; len(eval) != 0 {
		for _, dot := range eval {
			_ = dot
			if dot.Cmd == ":one" || dot.Cmd == ":many" {
				if dot.Ret.IsStruct() {
					lookup[dot.Ret.Type()] = dot
				}
			}
		}
	}

	return makeReaderString(lookup, v)
}

func Offset(v int) int {
	return v + 1
}

func ExecCommand(t TmplCtx, q Query) string {

	reader := ""
	if q.Ret.IsStruct() {
		reader = sdk.LowerTitle(q.Ret.Type()) + "Reader"
	} else {
		reader = fmt.Sprintf(`(fun r -> r.%s "%s")`, q.Ret.Typ.ReaderTyp, q.Ret.Typ.DbName)
	}

	if t.Configuration.Async {
		switch t.Settings.Engine {
		case "postgresql":
			switch q.Cmd {
			case ":one":
				return "executeRowAsync" + " " + reader
			case ":many":
				return "executeAsync" + " " + reader
			default:
				return "executeNonQueryAsync"
			}
		case "sqlite":
			switch q.Cmd {
			case ":one", ":many":
				return "executeAsync" + " " + reader
			default:
				return "executeNonQueryAsync"
			}
		}
	} else {
		switch t.Settings.Engine {
		case "postgresql":
			switch q.Cmd {
			case ":one":
				return "executeRow" + " " + reader
			case ":many":
				return "execute" + " " + reader
			default:
				return "executeNonQuery"
			}
		case "sqlite":
			switch q.Cmd {
			case ":one", ":many":
				return "execute" + " " + reader
			default:
				return "executeNonQuery"
			}

		}
	}
	return ""
}

func MakeConfig(req *plugin.Request) (Config, error) {

	var conf Config
	if len(req.PluginOptions) > 0 {
		if err := easyjson.Unmarshal(req.PluginOptions, &conf); err != nil {
			return conf, err
		}
	}
	return conf, nil
}

func Format(s string) string {
	// TODO: do more than just skip multiple blank lines, like maybe run fslint to format
	skipNextSpace := false
	var lines []string
	for _, l := range strings.Split(s, "\n") {
		isSpace := len(strings.TrimSpace(l)) == 0
		if !isSpace || !skipNextSpace {
			lines = append(lines, l)
		}
		skipNextSpace = isSpace
	}
	o := strings.Join(lines, "\n")
	o += "\n"
	return o
}
