package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var _db *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:root@/myDB")
	if err != nil {
		fmt.Println(err)
		return
	}
	_db = db

	fmt.Println("db init")
}

func CloseDB() {
	if _db != nil {
		fmt.Println("db closed")
		_db.Close()
	}
}

func Query(sqlstr string, result interface{}) error {

	err := _db.Ping()
	if err != nil {
		fmt.Println(err)
		return err
	}

	rows, err := _db.Query("Select id,data from think_data")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer rows.Close()

	var id int
	var data string

	for rows.Next() {
		err := rows.Scan(&id, &data)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("id=", id, "data=", data)

	}

	return nil
}
