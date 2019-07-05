package util

import (
	"fmt"
	test_fixtures "github.com/mchirico/activeIncident/test-fixtures"
	"github.com/mchirico/tlib/util"
	"golang.org/x/net/html"
	"log"
	"strings"
	"testing"
)

func TestBeging(t *testing.T) {
	s := test_fixtures.Page()
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}

func TestTag(t *testing.T) {
	defer util.NewTlib().ConstructDir()()

	result, err := Tag(test_fixtures.Page())
	if err != nil {
		t.FailNow()
	}

	strip(result[0])

}

func Test_LiveCheck(t *testing.T) {

	url := "https://webapp02.montcopa.org/eoc/cadinfo/livecad.asp"
	r, err := Get(url)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	result, err := Tag(r)
	if err != nil {
		t.FailNow()
	}

	strip(result[0])

}
