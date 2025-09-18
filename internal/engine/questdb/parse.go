package questdb

import (
	"io"

	"github.com/EnduIf/sqlc/internal/engine/postgresql"
	"github.com/EnduIf/sqlc/internal/source"
	"github.com/EnduIf/sqlc/internal/sql/ast"
)

// NewParser creates a new QuestDB parser
// QuestDB uses PostgreSQL-compatible SQL syntax, so we reuse the PostgreSQL parser
func NewParser() *Parser {
	return &Parser{
		pg: postgresql.NewParser(),
	}
}

type Parser struct {
	pg *postgresql.Parser
}

func (p *Parser) Parse(r io.Reader) ([]ast.Statement, error) {
	return p.pg.Parse(r)
}

func (p *Parser) CommentSyntax() source.CommentSyntax {
	return p.pg.CommentSyntax()
}

func (p *Parser) IsReservedKeyword(s string) bool {
	return p.pg.IsReservedKeyword(s)
}