package view

import (
	"log"
	"os"
)

type Static struct {
	DistDir string
}

func (s *Static) Generate() {

	if err := os.MkdirAll(s.DistDir, os.ModePerm); err != nil {
		log.Fatalf("Could not create destination directory: %v", err)
	}

	s.pageIndex()
}

func (s *Static) createFile(fileName string) *os.File {
	f, err := os.Create(s.DistDir + fileName)
	if err != nil {
		log.Fatalf("Could not create file: %v", err)
	}
	return f
}
