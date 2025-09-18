package questdb

import (
	"github.com/EnduIf/sqlc/internal/engine/postgresql"
	"github.com/EnduIf/sqlc/internal/sql/catalog"
)

// NewCatalog creates a new QuestDB catalog
// QuestDB uses PostgreSQL-compatible schema structure, so we reuse the PostgreSQL catalog
func NewCatalog() *catalog.Catalog {
	return postgresql.NewCatalog()
}