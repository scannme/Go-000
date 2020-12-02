package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

/*
	操作数据库时，如果遇到一个 <not found> 错误
	可以和其他错误一样，Wrap抛给上层，让交给上层决定是否对<not found>错误做处理
*/
func main() {
	if err := biz(); err != nil {
		fmt.Printf("%+v", err)
		if errors.Is(err, sql.ErrNoRows) {
			// 如果要根据特定的错误做特定的处理，使用errors.Is判断错误
			// ...
		} else {
			return
		}
	}
}

func biz() error {
	userID := 3
	if _, err := Dao(userID); err != nil {
		return err
	}

	// ...
	return nil
}

// User user information
type User struct {
	ID   int
	name string
	// ...
}

// Dao  find user by id
func Dao(id int) (string, error) {

	var name string

	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.QueryRow("SELECT ... WHERE id=?", id).Scan(&name)

	if sql.ErrNoRows == err {
		return "", errors.Wrapf(err, "id: %v", id)
	} else {
		return name, nil
	}
}
