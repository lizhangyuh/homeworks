package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client
var ctx context.Context

type TestCase struct {
	key       string
	valueSize int
	count     int
}

func main() {
	client = redis.NewClient(&redis.Options{
		Network:    "",
		Addr:       "127.0.0.1:6379",
		DB:         0,
		MaxRetries: 5,
		PoolSize:   128,
	})

	ctx = context.Background()

	tc_10b_10k := &TestCase{key: "10b_10k", valueSize: 10, count: 10000}
	tc_10b_10k.test_memory()
	tc_20b_10k := &TestCase{key: "20b_10k", valueSize: 20, count: 10000}
	tc_20b_10k.test_memory()
	tc_50b_10k := &TestCase{key: "50b_10k", valueSize: 50, count: 10000}
	tc_50b_10k.test_memory()
	tc_100b_10k := &TestCase{key: "100b_10k", valueSize: 100, count: 10000}
	tc_100b_10k.test_memory()
	tc_200b_10k := &TestCase{key: "200b_10k", valueSize: 200, count: 10000}
	tc_200b_10k.test_memory()
	tc_1kb_10k := &TestCase{key: "1kb_10k", valueSize: 1000, count: 10000}
	tc_1kb_10k.test_memory()
	tc_5kb_10k := &TestCase{key: "5kb_10k", valueSize: 5000, count: 10000}
	tc_5kb_10k.test_memory()

	tc_10b_50k := &TestCase{key: "10b_50k", valueSize: 10, count: 50000}
	tc_10b_50k.test_memory()
	tc_20b_50k := &TestCase{key: "20b_50k", valueSize: 20, count: 50000}
	tc_20b_50k.test_memory()
	tc_50b_50k := &TestCase{key: "50b_50k", valueSize: 50, count: 50000}
	tc_50b_50k.test_memory()
	tc_100b_50k := &TestCase{key: "100b_50k", valueSize: 100, count: 50000}
	tc_100b_50k.test_memory()
	tc_200b_50k := &TestCase{key: "200b_50k", valueSize: 200, count: 50000}
	tc_200b_50k.test_memory()
	tc_1kb_50k := &TestCase{key: "1kb_50k", valueSize: 1000, count: 50000}
	tc_1kb_50k.test_memory()
	tc_5kb_50k := &TestCase{key: "5kb_50k", valueSize: 5000, count: 50000}
	tc_5kb_50k.test_memory()

}

func (tc *TestCase) test_memory() {
	client.FlushAll(ctx)
	before := info_memory()
	tc.write()
	after := info_memory()

	result := fmt.Sprintf("before: \n %s \n after: \n %s \n", before, after)
	os.WriteFile(fmt.Sprintf("./%s", tc.key), []byte(result), 0644)
}

func (tc *TestCase) write() error {
	for i := 0; i < tc.count; i++ {
		status := client.Set(ctx, fmt.Sprintf("%s_%d", tc.key, i), generate_value(tc.valueSize), -1)
		if status.Err() != nil {
			return status.Err()
		}
	}

	return nil
}

func generate_value(size int) string {
	val := make([]byte, size)
	for i := 0; i < size; i++ {
		val[i] = 'r'
	}

	return string(val)
}

func info_memory() string {
	cmd := client.Info(ctx, "memory")
	return cmd.Val()
}
