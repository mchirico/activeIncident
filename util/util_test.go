/*
https://webapp02.montcopa.org/eoc/cadinfo/livecad.asp?print=yes

 */


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

func TestBegin(t *testing.T) {
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

	result, _, err := Tag(test_fixtures.Page())
	if err != nil {
		t.FailNow()
	}

	strip(result[0])

}

func TestGetTable(t *testing.T) {
	r := test_fixtures.Table()
	result, err := GetTable(r)
	if err != nil {
		t.FailNow()
	}
	fmt.Printf("%v\n", result)
}

func Test_LiveCheck(t *testing.T) {

	url := "https://webapp02.montcopa.org/eoc/cadinfo/livecad.asp"
	r, err := Get(url)
	if err != nil {
		t.Fatalf("err: %s\n", err)
	}

	result, link, err := Tag(r)
	if err != nil {
		t.FailNow()
	}

	strip(result[0])

	for _, l := range link {
		r, err = Get(GetDetail(l))
		if err != nil {
			t.Fatalf("err: %s\n", err)
		}

		result, err = GetTable(r)
		fmt.Printf("%v\n", result)

	}

}


func TestGetBuildDB(t *testing.T) {

	c,a,err := BuildDb()
	if err != nil {
		t.FailNow()
	}
	for i,m := range c {
		for k,v := range m {
			fmt.Printf("%v: %v\n",k,v)
		}
		fmt.Printf("Status: %v\n\n",a[i])
	}


}

func TestShow(t *testing.T) {
	Show()
}

func TestGetTableV2(t *testing.T) {

	a,_ := GetTableV2(test_fixtures.Detail())
	for _,v :=range a {
		fmt.Printf("%v\n",v)
	}
}