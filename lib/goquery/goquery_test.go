package goquery

import (
	"fmt"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestFind(t *testing.T) {
	html := `<html><body>
	<div id="main">
	  <div></div>
	</div>
	</body></html>`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		panic(err)
	}
	nodes := doc.Find("#main").Nodes
	for _, node := range nodes {
		fmt.Printf("%+v\n", node)
	}
}
