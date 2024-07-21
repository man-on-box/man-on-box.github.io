package view

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/yuin/goldmark"
)

type Static struct {
	DistDir string
}

func (s *Static) Generate() {
	tmpl := parseTemplates()
	s.copyPublicDir()

	s.pageIndex(tmpl)
	s.pageAbout(tmpl)
}

func parseTemplates() *template.Template {
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"currentYear": currentYear,
	}).ParseGlob("./view/templates/**/*.html")
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}
	return tmpl
}

func (s *Static) copyPublicDir() {
	if err := copyDir("public", s.DistDir); err != nil {
		log.Fatalf("Could not copy public directory: %v", err)
	}
}

func (s *Static) createFile(fileName string) *os.File {
	f, err := os.Create(s.DistDir + fileName)
	if err != nil {
		log.Fatalf("Could not create file: %v", err)
	}
	return f
}

func mdFileToHTML(file string) template.HTML {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading markdown file: %v", err)
	}
	var buf bytes.Buffer
	if err := goldmark.Convert(f, &buf); err != nil {
		log.Fatalf("Error converting markdown to HTML: %v", err)
	}
	return template.HTML(buf.String())
}

func mdToHTML(md string) template.HTML {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(md), &buf); err != nil {
		log.Fatalf("Error converting markdown to HTML: %v", err)
	}
	return template.HTML(buf.String())
}

func currentYear() int {
	return time.Now().Year()
}

func copyFile(src string, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

func copyDir(src string, dst string) error {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		return err
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}
