package pages

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (p *Pages) serve(port string) {
	fmt.Println("Starting dev server on port:", port)
	var rootHandler http.HandlerFunc

	for _, page := range *p.pages {
		fmt.Println("Serving page: ", page.filePath)

		handler := func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("%s: %s\n", r.Method, r.URL)
			page.render(w)
		}

		http.HandleFunc(page.filePath, handler)
		http.HandleFunc(strings.TrimSuffix(page.filePath, ".html"), handler)

		if page.filePath == "/index.html" {
			rootHandler = handler
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			rootHandler(w, r)
		} else {
			http.ServeFile(w, r, "public"+r.URL.Path)
		}
	})

	err := http.ListenAndServe("localhost:"+port, nil)
	if err != nil {
		log.Fatalf("Could not start dev server: %v", err)
	}
}
