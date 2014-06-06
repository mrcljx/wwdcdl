package main

import (
	"code.google.com/p/go-html-transform/h5"
	css "code.google.com/p/go-html-transform/css/selector"
	"code.google.com/p/go.net/html"
	"io/ioutil"
	"net/http"
)

func FindNodes(node *html.Node, selector string) []*html.Node {
	parsedSelector, _ := css.Selector(selector)
	return parsedSelector.Find(node)
}

func FindNode(node *html.Node, selector string) *html.Node {
	return FindNodes(node, selector)[0]
}

func ParseHtmlAt(url string) (node *html.Node, err error) {
	resp, err := http.Get(url)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	tree, err := h5.NewFromString(string(body))

	if err != nil {
		return
	}

	node = tree.Top()
	return
}

func GetAttrValue(node *html.Node, attrKey string) string {
	for _, attr := range node.Attr {
		if attr.Key == attrKey {
			return attr.Val
		}
	}

	return ""
}

func InnerHtml(node *html.Node) string {
	return h5.RenderNodesToString(h5.Children(node))
}
