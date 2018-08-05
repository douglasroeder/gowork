package tests

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// SQLite dbEngine
type SQLite struct {
	db *gorm.DB
}

// NewSQLiteEngine returns SQLite engine that knows how to truncate a table
func NewSQLiteEngine(db *gorm.DB) *SQLite {
	return &SQLite{
		db: db,
	}
}

// Truncate cleans up table
func (sqlite *SQLite) Truncate(table string) error {
	tx := sqlite.db.Begin()

	cmds := []string{
		fmt.Sprintf("DELETE FROM %s", table),
	}

	for _, cmd := range cmds {
		tx.Raw(cmd)
	}

	tx.Commit()
	return nil
}

// Close ends db connection
func (sqlite *SQLite) Close() error {
	return sqlite.db.Close()
}
