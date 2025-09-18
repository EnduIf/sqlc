package golang

import (
	"github.com/EnduIf/sqlc/internal/codegen/golang/opts"
	"github.com/EnduIf/sqlc/internal/codegen/sdk"
	"github.com/EnduIf/sqlc/internal/plugin"
)

func questdbType(req *plugin.GenerateRequest, options *opts.Options, col *plugin.Column) string {
	columnType := sdk.DataType(col.Type)
	notNull := col.NotNull || col.IsArray
	
	// QuestDB specific type mappings
	switch columnType {
	case "symbol":
		// QuestDB's symbol type is essentially a string with optimized storage
		if notNull {
			return "string"
		}
		return "sql.NullString"
		
	case "geohash":
		// QuestDB's geohash type for location data
		if notNull {
			return "string"
		}
		return "sql.NullString"
	}
	
	// For most other types, QuestDB is PostgreSQL-compatible
	// so we reuse the PostgreSQL type mappings
	return postgresType(req, options, col)
}