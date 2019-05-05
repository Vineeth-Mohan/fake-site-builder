package main

import (
	"flag"
	"graph"
	"log"
	"net/http"
)

func main() {
	templatePath := flag.String("template", "./seed-pages", "Path to seed HTML pages")
	dumpPath := flag.String("dumpPath", "site", "Directory where fabricated HTML would be stored")
	numOfPages := flag.Int("pageCount", 100, "Number of pages to generate")
	flag.Parse()

	graph := graph.Graph{}
	dumpFolder := *dumpPath
	graph.BuildGraph(
		*numOfPages,
		*templatePath,
		dumpFolder)

	fs := http.FileServer(http.Dir(dumpFolder))
	http.Handle("/", fs)

	log.Println("Bringing Static Server Up ...")
	http.ListenAndServe(":3000", nil)

}
