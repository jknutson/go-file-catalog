package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v4"
)

var baseDir = "/Users/john.knutson/Downloads/3BV9NgkoDGWk"

func main() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	err = filepath.Walk(baseDir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			f, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()

			if !info.IsDir() {
				h := sha256.New()
				if _, err := io.Copy(h, f); err != nil {
					log.Fatal(err)
				}
				shaSum := fmt.Sprintf("%x", h.Sum(nil))
				fmt.Println(path, filepath.Base(path), info.Size(), shaSum)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

/* batch query/insert; if `i % batchSize == 0` then SendBatch
batch := &pgx.Batch{}
batch.Queue("insert into ledger(description, amount) values($1, $2)", "q1", 1)
batch.Queue("insert into ledger(description, amount) values($1, $2)", "q2", 2)
br := conn.SendBatch(context.Background(), batch)
*/
