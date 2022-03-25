//Q: 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	user, err := FetchUserByName("Jerry")

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// 处理没有查询到数据的情况
			fmt.Printf("%+v\n", err)
		}

		return
	}

	fmt.Printf("got user %v\n", user.Name)
}
