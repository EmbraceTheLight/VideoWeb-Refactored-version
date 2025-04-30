package dbutil

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SqlTxToGormDB is a helper function.
// It returns a *gorm.DB from *sql.Tx.
func SqlTxToGormDB(tx *sql.Tx) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: tx,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
