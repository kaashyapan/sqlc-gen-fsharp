package core

import (
	"log"
	"strings"

	plugin "github.com/tabbed/sqlc-go/codegen"
	"github.com/tabbed/sqlc-go/sdk"
)

// https://www.npgsql.org/doc/types/basic.html
// https://www.postgresql.org/docs/current/datatype-numeric.html
// returns f# type, reader type, library type, enum flag
func postgresType(req *plugin.CodeGenRequest, col *plugin.Column) (string, string, string, bool) {
	columnType := strings.ToLower(sdk.DataType(col.Type))
	config, _ := MakeConfig(req)

	if config.TypeAffinity {
		switch columnType {
		case "serial", "serial4", "pg_catalog.serial4", "smallserial", "serial2", "pg_catalog.serial2", "integer", "int", "int4", "pg_catalog.int4", "smallint", "int2", "pg_catalog.int2":
			if col.NotNull {
				return "int", "int", "int", false
			} else {
				return "int option", "intOrNone", "intOrNone", false
			}
		case "float", "double", "double precision", "float8", "pg_catalog.float8", "real", "float4", "pg_catalog.float4":

			if col.NotNull {
				return "double", "double", "double", false
			} else {
				return "double option", "doubleOrNone", "doubleOrNone", false
			}

		}
	}

	switch columnType {
	case "serial", "serial4", "pg_catalog.serial4":
		if col.NotNull {
			return "int", "int", "int", false
		} else {
			return "int option", "intOrNone", "intOrNone", false
		}
	case "bigserial", "serial8", "pg_catalog.serial8":
		if col.NotNull {
			return "int64", "int64", "int64", false
		} else {
			return "int64 option", "int64OrNone", "int64OrNone", false
		}
	case "smallserial", "serial2", "pg_catalog.serial2":
		if col.NotNull {
			return "int16", "int16", "int16", false
		} else {
			return "int16 option", "int16OrNone", "int16OrNone", false
		}
	case "integer", "int", "int4", "pg_catalog.int4":
		if col.NotNull {
			return "int", "int", "int", false
		} else {
			return "int option", "intOrNone", "intOrNone", false
		}
	case "bigint", "int8", "pg_catalog.int8":
		if col.NotNull {
			return "int64", "int64", "int64", false
		} else {
			return "int64 option", "int64OrNone", "int64OrNone", false
		}
	case "smallint", "int2", "pg_catalog.int2":
		if col.NotNull {
			return "int16", "int16", "int16", false
		} else {
			return "int16 option", "int16OrNone", "int16OrNone", false
		}
	case "float", "double", "double precision", "float8", "pg_catalog.float8":
		if col.NotNull {
			return "double", "double", "double", false
		} else {
			return "double option", "doubleOrNone", "doubleOrNone", false
		}
	case "real", "float4", "pg_catalog.float4":

		if col.NotNull {
			return "float32", "real", "real", false
		} else {
			return "float32 option", "realOrNone", "realOrNone", false
		}
	case "numeric", "money", "pg_catalog.numeric":
		if col.NotNull {
			return "decimal", "decimal", "decimal", false
		} else {
			return "decimal option", "decimalOrNone", "decimalOrNone", false
		}
	case "boolean", "bool", "pg_catalog.bool":
		if col.NotNull {
			return "bool", "bool", "bool", false
		} else {
			return "bool option", "boolOrNone", "boolOrNone", false
		}
	case "jsonb", "json":
		if col.NotNull {
			return "string", "string", "jsonb", false
		} else {
			return "string option", "stringOrNone", "jsonbOrNone", false
		}
	case "bytea", "blob", "pg_catalog.bytea":
		if col.NotNull {
			return "byte[]", "bytea", "bytea", false
		} else {
			return "byte[] option", "byteaOrNone", "byteaOrNone", false
		}
	case "date":
		if col.NotNull {
			return "DateOnly", "dateOnly", "date", false
		} else {
			return "DateOnly option", "dateOnlyOrNone", "dateOrNone", false
		}
	case "pg_catalog.time":
		if col.NotNull {
			return "TimeSpan", "interval", "interval", false
		} else {
			return "TimeSpan option", "intervalOrNone", "intervalOrNone", false
		}
	case "pg_catalog.timestamp":
		if col.NotNull {
			return "DateTime", "dateTime", "timestamp", false
		} else {
			return "DateTime option", "dateTimeOrNone", "timestampOrNone", false
		}
	case "pg_catalog.timestamptz", "timestamptz", "pg_catalog.timetz":
		// TODO
		if col.NotNull {
			return "DateTimeOffset", "datetimeOffset", "timestamptz", false
		} else {
			return "DateTimeOffset option", "datetimeOffsetOrNone", "timestamptzOrNone", false
		}
	case "text":
		if col.NotNull {
			return "string", "text", "text", false
		} else {
			return "string option", "textOrNone", "textOrNone", false
		}

	case "pg_catalog.varchar", "pg_catalog.bpchar", "string":
		if col.NotNull {
			return "string", "string", "string", false
		} else {
			return "string option", "stringOrNone", "stringOrNone", false
		}
	case "uuid":
		if col.NotNull {
			return "Guid", "uuid", "uuid", false
		} else {
			return "Guid option", "uuidOrNone", "uuidOrNone", false
		}

	case "point":
		if col.NotNull {
			return "NpgsqlPoint", "point", "point", false
		} else {
			return "NpgsqlPoint option", "pointOrNone", "pointOrNone", false
		}

	case "void", "null", "NULL":
		// TODO
		// A void value always returns NULL. Since there is no built-in NULL
		// value into the SQL package, we'll use sql.NullBool
		return "System.Nullable", "dbNull", "dbNull", false

	default:
		// TODO Enums
		log.Printf("unknown PostgreSQL type: %s\n", columnType)
		return columnType, columnType + "_unhandled_report_issue", columnType + "_unhandled_report_issue", false
	}
}
