package graph

import (
	"Helper"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// Node :- Represents Single Page
type Node struct {
	URL           string
	doc           goquery.Document
	outboundLinks []string
}

// Graph :- Maintain the entire graph state of the site
type Graph struct {
	URLConnections map[string][]Node
	Nodes          []Node
}

//BuildGraph :- Builds the graph representinf the site
func (graph *Graph) BuildGraph(count int, templateHTMLPath string, dumpFolder string) error {
	graph.Nodes = []Node{}
	pageDocs, error := graph.ReadPages(count, templateHTMLPath)
	graph.CreateAllPages(pageDocs)
	graph.createOutboundLinks()
	graph.dumpSite(dumpFolder)
	return error
}

func (graph *Graph) ReadPages(count int, pathToTemplates string) ([]goquery.Document, error) {
	var docs []goquery.Document
	fmt.Println("Folder -> ", pathToTemplates)
	err := filepath.Walk(pathToTemplates,
		func(path string, info os.FileInfo, err error) error {
			fmt.Println("Path -> ", info.Name())
			if info.IsDir() {
				return nil
			}
			doc, err := graph.createDocument(path)
			if err != nil {
				return err
			}
			docs = append(docs, *doc)
			return err
		})
	docs = graph.correctDocsLength(count, docs)
	if err != nil {
		return nil, err
	}
	return docs, nil
}
func (graph *Graph) correctDocsLength(count int, docs []goquery.Document) []goquery.Document {
	if len(docs) > count {
		docs = docs[1:count]
	}
	if len(docs) < count {
		docsToCreate := count - len(docs)
		for i := 0; i < docsToCreate; i++ {
			index := i % len(docs)
			docs = append(docs, docs[index])
		}
	}
	return docs
}
func (graph *Graph) createDocument(path string) (*goquery.Document, error) {
	body, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func RemoveNode(root_node *html.Node, remove_me *html.Node) {
	found_node := false
	check_nodes := make(map[int]*html.Node)
	i := 0

	// loop through siblings
	for n := root_node.FirstChild; n != nil; n = n.NextSibling {
		if n == remove_me {
			found_node = true
			n.Parent.RemoveChild(n)
		}

		check_nodes[i] = n
		i++
	}

	// check if removing node is found
	// if yes no need to check childs returning
	// if no continue loop through childs and so on
	if found_node == false {
		for _, item := range check_nodes {
			RemoveNode(item, remove_me)
		}
	}
}

func (graph *Graph) CreateAllPages(pageDocs []goquery.Document) {
	for idx, pageDoc := range pageDocs {
		page := "index"
		if idx != 0 {
			page = Helper.CreateRandomString(10)
		}
		pageURL := "/" + page + ".html"
		graph.Nodes = append(graph.Nodes, Node{URL: pageURL, doc: pageDoc})
	}
}

func (graph *Graph) createOutboundLinks() error {
	for _, node := range graph.Nodes {
		graph.buildPage(&node.doc)
	}
	return nil
}

func (graph *Graph) buildPage(doc *goquery.Document) {
	doc.Find("script").Remove()
	doc.Find("a").Each(func(i int, anchorElement *goquery.Selection) {
		nodeIdx := Helper.CreateRandomInteger(len(graph.Nodes))
		outboundURL := graph.Nodes[nodeIdx].URL
		anchorElement.SetAttr("href", outboundURL)
	})
	// html, _ := doc.Html()
	// fmt.Println(html)
}

func (graph *Graph) dumpSite(staticFolder string) error {
	err := Helper.CreateFolderIfNotExist(staticFolder)
	if err != nil {
		return err
	}
	for _, node := range graph.Nodes {
		html, err := node.doc.Html()
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(staticFolder+node.URL,
			[]byte(html),
			0644)
		if err != nil {
			return err
		}
	}
	return nil
}
