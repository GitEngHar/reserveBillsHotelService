package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	dbUser     = "root"
	dbPassword = "password"
	dbHost     = "localhost"
	dbPort     = "3306"
	dbName     = "hotel_db"
)

func NewMySQL() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	// コネクションプールの設定
	db.SetMaxOpenConns(5)                   // 最大接続数
	db.SetMaxIdleConns(2)                   // アイドル接続数
	db.SetConnMaxLifetime(30 * time.Minute) // 最大接続時間

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ MySQL 接続成功！")
	return db, nil
}
