package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"io/fs"
	"io/ioutil"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

type result struct {
	path string
	sum  [16]byte
}

func main() {
	m, err := MD5All(context.Background(), ".")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	for k, v := range m {
		fmt.Printf("%s: \t%x\n", k, v)
	}
}

func MD5All(ctx context.Context, root string) (map[string][16]byte, error) {
	g, ctx := errgroup.WithContext(ctx)

	paths := make(chan string)

	g.Go(func() error {
		defer close(paths)

		filepath.Walk(root, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			select {
			case paths <- path:
			case <-ctx.Done():
				return ctx.Err()
			}

			return nil
		})

		return nil
	})

	c := make(chan result)
	const numDigesters = 20
	for i := 0; i < numDigesters; i++ {
		g.Go(func() error {
			for path := range paths {
				data, err := ioutil.ReadFile(path)
				if err != nil {
					return err
				}

				select {
				case c <- result{path, md5.Sum(data)}:
				case <-ctx.Done():
					return ctx.Err()
				}
			}
			return nil
		})
	}

	go func() {
		g.Wait()
		close(c)
	}()

	m := make(map[string][16]byte)

	for r := range c {
		m[r.path] = r.sum
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return m, nil
}
