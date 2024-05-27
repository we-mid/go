package db

import (
	"database/sql"
)

func NewDB(driver, dsn, migrationDir string) (*sql.DB, error) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	// 验证连接
	if err := db.Ping(); err != nil {
		return nil, err
	}
	// 启用外键约束 (sqlite)
	// _, err = db.Exec("PRAGMA foreign_keys=ON;")
	// if err != nil {
	// 	return nil, err
	// }
	// 数据库迁移逻辑
	if err := dbMigrate(db, migrationDir); err != nil {
		return nil, err
	}
	return db, nil
}
