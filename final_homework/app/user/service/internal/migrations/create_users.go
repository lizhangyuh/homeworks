package main

import (
	"flag"
	"fmt"

	"nvm.com/mrc-api-go/app/user/service/internal/conf"
	"nvm.com/mrc-api-go/app/user/service/internal/data"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"gorm.io/gorm"
)

var flagconf string

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	db := getDB()

	db.AutoMigrate(&data.User{})
	fmt.Println("migrations done!")
}

func getDB() *gorm.DB {
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	return data.NewDB(bc.Data)
}
