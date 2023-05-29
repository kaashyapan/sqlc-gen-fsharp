package core

import (
	plugin "github.com/tabbed/sqlc-go/codegen"
	"github.com/tabbed/sqlc-go/sdk"
)

func mysqlType(req *plugin.CodeGenRequest, col *plugin.Column) (string, string, string, bool) {
	columnType := sdk.DataType(col.Type)

	switch columnType {

	case "varchar", "text", "char", "tinytext", "mediumtext", "longtext":
		if col.NotNull {
			return "string", "string", "string", false
		} else {
			return "string option", "string", "string", false

		}
	case "int", "integer", "smallint", "mediumint", "year":
		if col.NotNull {
			return "int", "int", "int", false
		} else {
			return "int option", "int", "int", false

		}

	case "bigint":
		if col.NotNull {
			return "int64", "int64", "int64", false
		} else {
			return "int64 option", "int64", "int64", false

		}

	case "blob", "binary", "varbinary", "tinyblob", "mediumblob", "longblob":
		if col.NotNull {
			return "byte[]", "byte[]", "byte[]", false
		} else {
			return "byte[] option", "byte[]", "byte[]", false

		}

	case "double", "double precision":
		if col.NotNull {
			return "double", "double", "double", false
		} else {
			return "double option", "double", "double", false

		}

	case "real":
		if col.NotNull {
			return "real", "real", "real", false
		} else {
			return "float option", "real", "real", false

		}

	case "decimal", "dec", "fixed":
		if col.NotNull {
			return "decimal", "decimal", "decimal", false
		} else {
			return "decimal option", "decimal", "decimal", false

		}

	case "date", "datetime", "time":
		if col.NotNull {
			return "DateTime", "DateTime", "DateTime", false
		} else {
			return "DateTime option", "DateTime", "DateTime", false

		}

	case "timestamp":
		if col.NotNull {
			return "DateTimeOffset", "DateTimeOffset", "DateTimeOffset", false
		} else {
			return "DateTimeOffset option", "DateTimeOffset", "DateTimeOffset", false

		}

	case "boolean", "bool", "tinyint":
		if col.NotNull {
			return "bool", "bool", "bool", false
		} else {
			return "bool option", "bool", "bool", false

		}

	case "json":
		if col.NotNull {
			return "string", "string", "string", false
		} else {
			return "string option", "string", "string", false

		}

	case "any":
		return "obj", "obj", "obj", false

	default:
		return columnType, columnType + "_unhandled_report_issue", columnType + "_unhandled_report_issue", false

	}
}
