package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var baseDir = "."

func main() {
	err := filepath.Walk(baseDir,
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
				fmt.Println(path, info.Size(), shaSum)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
