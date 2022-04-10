package data

import (
	"user/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewUserRepo)

// Data .
type Data struct {
	db *gorm.DB
	// rdb *redis.Client
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

func NewRedis(c *conf.Data) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		DB:           int(c.Redis.Db),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})

	rdb.AddHook(redisotel.TracingHook{})

	if err := rdb.Close(); err != nil {
		log.Error(err)
	}

	return rdb
}
