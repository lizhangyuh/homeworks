package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"nvm.com/mrc-api-go/app/post/service/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewPostRepo)

// Data .
type Data struct {
	db *gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{db: db}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Errorf("failed opening connection to db: %v", err)
		panic("failed to connect db")
	}

	return db
}
