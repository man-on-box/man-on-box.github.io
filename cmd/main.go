package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/man-on-box/man-on-box.github.io/view"
)

const (
	envVar  = "APP_ENV"
	distDir = "dist"
)

func main() {
	env := os.Getenv(envVar)

	if err := os.MkdirAll(distDir, os.ModePerm); err != nil {
		log.Fatalf("Could not create destination directory: %v", err)
	}

	f, err := os.Create(distDir + "/index.html")
	if err != nil {
		log.Fatalf("Could not create file: %v", err)
	}

	component := view.Index()
	err = component.Render(context.Background(), f)
	if err != nil {
		log.Fatalf("Could not render component: %v", err)
	}

	if env == "dev" {
		startDevServer()
	}
}

func startDevServer() {
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)

	fmt.Println("Starting dev server on http://localhost:8080")

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalf("Could not start dev server: %v", err)
	}
}
