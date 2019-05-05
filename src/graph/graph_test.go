package graph

import (
	"fmt"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

var (
	html = `
	<html><body><div class="outter-class">
        <h1 class="inner-class">
			The string I need
			<script>
				alert("something 1")
			</script>
	        <span class="other-class" >Some value I don't need</span>
	        <span class="other-class2" title="sometitle"></span>
        </h1>
        <div class="other-class3">
			<h3>Some heading i don't need</h3>			
			<script>
				alert("something 2")
			</script>

		</div>
		<script>
			alert("something 3")
		</script>

    </div></body></html>`
)

func TestReadPages(t *testing.T) {
	graph := Graph{}
	docs, error := graph.ReadPages(100,
		"/Users/vineeth/git/site-generator/test-data")
	if error != nil {
		t.Errorf("Error file reading file itself")
	}
	var titles []string
	for _, doc := range docs {
		doc.Find("title").Each(func(i int, titleElement *goquery.Selection) {
			title := titleElement.Text()
			titles = append(titles, title)
			fmt.Println("Title ->", title)
		})
	}
	titlePresentForAll := true
	for _, title := range titles {
		if len(title) == 0 {
			titlePresentForAll = false
		}
	}
	if !titlePresentForAll {
		t.Errorf("Title not present for all pages")
	}
}
