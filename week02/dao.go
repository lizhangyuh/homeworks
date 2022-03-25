package main

import (
	"database/sql"
	"errors"

	xerrors "github.com/pkg/errors"
)

type User struct {
	Name string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:mfitness@tcp(127.0.0.1:3306)/mfitness?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
}

func FetchUserByName(name string) (*User, error) {
	return FetchUser("SELECT name FROM users WHERE name = ?", name)
}

func FetchUser(query string, params ...interface{}) (*User, error) {
	row := Db.QueryRow(query, params...)

	var user User
	if err := row.Scan(&user.Name); err != nil {
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			// 包装错误返回给上层
			// 在DAO中查询不到数据不能算是一个真正意义上的错误，应该把错误返回给上层，交由上层业务逻辑代码处理查询数据为空的情况
			err = xerrors.Wrapf(err, "user dao: not found, %v", query)
			return nil, err
		}
	}

	return &user, nil

}
