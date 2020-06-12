package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type miracle struct {
	id         int
	name       string
	memoryCell int
	useTime    int
}

// Connectsql  根据数据库信息连接数据库
func Connectsql(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("can not connect %v\n", err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("failed %v\n", err)
		return nil
	}
	fmt.Println("connected")
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(3)
	return db
}

// Insert 插入数据
func Insert(db *sql.DB) {
	var (
		sqlStr      string   = "INSERT INTO miracles(id, name, memoryCell, useTime) VALUES(?,?,?,?);"
		nameSlice   []string = []string{"Great lighting Spear", "Sunlight Spear", "Sunlight Blade", "Darkmoon Blade", "Wrath of The Gods"}
		memorySlice []int    = []int{1, 2, 1, 1, 1}
		useSlice    []int    = []int{10, 5, 1, 1, 3}
	)
	for i := 0; i < 5; i++ {
		_, err := db.Exec(sqlStr, i+2, nameSlice[i], memorySlice[i], useSlice[i])
		if err != nil {
			fmt.Printf("can not insert %v\n", err)
			return
		}
	}
	fmt.Println("mission complete")
}

// Delete 删除数据
func Delete(db *sql.DB) {
	var sqlStr string = "DELETE FROM miracles WHERE id=?"
	ret, err := db.Exec(sqlStr, 6)
	if err != nil {
		fmt.Printf("can not delete %v\n", err)
		return
	}
	num, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("failed %v\n", err)
		return
	}
	fmt.Printf("Affected %v rows", num)
}

// Update 更改数据
func Update(db *sql.DB) {
	var sqlStr string = "UPDATE miracles SET useTime=? WHERE id=?"
	ret, err := db.Exec(sqlStr, 5, 6)
	if err != nil {
		fmt.Printf("can not update %v\n", err)
		return
	}
	num, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("failed %v\n", err)
		return
	}
	fmt.Printf("affected %v rows", num)
}

// QueryRows 数据库查询
func QueryRows(db *sql.DB) {
	var sqlStr string = "SELECT id,name,memoryCell,useTime FROM miracles WHERE id>?;"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("can not query %v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var m miracle
		err = rows.Scan(&m.id, &m.name, &m.memoryCell, &m.useTime)
		if err != nil {
			fmt.Printf("can not scan %v\n", err)
			return
		}
		fmt.Printf("id: %v, name: %v, memory: %v, usetime: %v.\n", m.id, m.name, m.memoryCell, m.useTime)
	}
}

func main() {
	var dsn string = "username:password@tcp(127.0.0.1:3306)/database"
	db := Connectsql(dsn)
	QueryRows(db)
}
