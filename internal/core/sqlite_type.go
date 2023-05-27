package core

import (
	"strings"

	plugin "github.com/tabbed/sqlc-go/codegen"
	"github.com/tabbed/sqlc-go/sdk"
)

// https://learn.microsoft.com/en-us/dotnet/standard/data/sqlite/types
func sqliteType(req *plugin.CodeGenRequest, col *plugin.Column) (string, string, string, bool) {

	columnType := strings.ToLower(sdk.DataType(col.Type))
	notNull := col.NotNull || col.IsArray

	switch columnType {

	case "int", "integer", "tinyint", "smallint", "mediumint", "bigint", "unsignedbigint", "int2", "int8":
		if notNull {
			return "int", "int", "int", false
		} else {
			return "int option", "intOrNone", "intOrNone", false
		}
	case "blob":
		if notNull {
			return "byte[]", "bytes", "bytes", false
		} else {
			return "byte[] option", "bytesOrNone", "bytesOrNone", false
		}
	case "real", "double", "doubleprecision", "float":
		if notNull {
			return "double", "double", "double", false
		} else {
			return "double option", "doubleOrNone", "doubleOrNone", false
		}
	case "boolean", "bool":
		if col.NotNull {
			return "bool", "bool", "bool", false
		} else {
			return "bool option", "boolOrNone", "boolOrNone", false
		}

	case "date", "datetime":
		if notNull {
			return "DateTime", "dateTime", "dateTime", false
		} else {
			return "DateTime option", "dateTimeOrNone", "dateTimeOrNone", false
		}
	case "timestamp":
		if notNull {
			return "DateTimeOffset", "dateTimeOffset", "dateTimeOffset", false
		} else {
			return "DateTimeOffset option", "dateTimeOffsetOrNone", "dateTimeOffsetOrNone", false
		}

	}

	switch {

	case strings.HasPrefix(columnType, "character"),
		strings.HasPrefix(columnType, "varchar"),
		strings.HasPrefix(columnType, "varyingcharacter"),
		strings.HasPrefix(columnType, "nchar"),
		strings.HasPrefix(columnType, "nativecharacter"),
		strings.HasPrefix(columnType, "nvarchar"),
		columnType == "text",
		columnType == "clob":
		if notNull {
			return "string", "string", "string", false
		} else {
			return "string option", "stringOrNone", "stringOrNone", false
		}

	case strings.HasPrefix(columnType, "decimal"), columnType == "numeric":
		if notNull {
			return "decimal", "decimal", "decimal", false
		} else {
			return "decimal option", "decimalOrNone", "decimalOrNone", false
		}

	default:
		return columnType, "unhandled_report_issue", "unhandled_report_issue", false

	}

}
