package helpers

import (
	"bufio"
	"bytes"
	"html/template"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/frontmatter"
)

var mdParser = goldmark.New(
	goldmark.WithExtensions(
		&frontmatter.Extender{},
	),
)

type ParsedMd struct {
	Html      template.HTML
	WordCount int
}

func MdFileToHTML(file string, metadata any) ParsedMd {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Error reading markdown file: %v", err)
	}
	var buf bytes.Buffer
	ctx := parser.NewContext()
	if err := mdParser.Convert(f, &buf, parser.WithContext(ctx)); err != nil {
		log.Fatalf("Error converting markdown to HTML: %v", err)
	}
	if metadata != nil {
		d := frontmatter.Get(ctx)
		if err := d.Decode(metadata); err != nil {
			log.Fatalf("Error decoding frontmatter data: %v", err)
		}
	}
	return ParsedMd{
		Html:      template.HTML(buf.String()),
		WordCount: getWordCount(string(f)),
	}
}

func getWordCount(s string) int {
	// Remove frontmatter from markdown content
	frontmatterRegex := regexp.MustCompile(`(?s)^---.*?---\s*|(?s)^\+\+\+.*?\+\+\+\s*`)
	s = frontmatterRegex.ReplaceAllString(s, "")
	// Naive remove html tags
	htmlTagRegex := regexp.MustCompile(`<.*?>`)
	s = htmlTagRegex.ReplaceAllString(s, "")

	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)

	count := 0
	for scanner.Scan() {
		count++
	}
	return count
}
