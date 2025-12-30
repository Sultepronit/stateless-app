package grabber

import (
	"bytes"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Tag struct {
	name      string
	attrName  string
	attrValue string
}

func checkAttribute(n *html.Node, attr string, value string) bool {
	for _, a := range n.Attr {
		if a.Key == attr && a.Val == value {
			return true
		}
	}

	return false
}

func checkNode(n *html.Node, tag Tag) bool {
	if n.Data == tag.name {
		if tag.attrName == "" {
			return true
		}
		return checkAttribute(n, tag.attrName, tag.attrValue)
	}

	return false
}

func collectNodes(n *html.Node, tag Tag) []*html.Node {
	re := make([]*html.Node, 0, 5)

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.ElementNode {
				if checkNode(c, tag) {
					re = append(re, c)
				} else {
					traverse(c)
				}
			}
		}
	}
	traverse(n)

	return re
}

func findNode(n *html.Node, tag string) *html.Node {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode {
			continue
		}

		if c.Data == tag {
			return c
		} else {
			return findNode(c, tag)
		}
	}

	return nil
}

func getTextContent(n *html.Node) string {
	var sb strings.Builder

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			switch c.Type {
			case html.TextNode:
				sb.WriteString(c.Data)
			case html.ElementNode:
				traverse(c)
			}
		}
	}
	traverse(n)

	return sb.String()
}

func unwrapTags(n *html.Node, tag string) {
	for c := n.FirstChild; c != nil; {
		if c.Type != html.ElementNode {
			c = c.NextSibling
			continue
		}

		next := c.NextSibling

		if c.Data == tag {
			for cc := c.FirstChild; cc != nil; {
				nx := cc.NextSibling
				c.RemoveChild(cc) // you must remove to insert in the other place!
				n.InsertBefore(cc, c)
				cc = nx
			}

			n.RemoveChild(c)
		} else {
			unwrapTags(c, tag)
		}

		c = next
	}
}

func nodeToHtml(n *html.Node) string {
	var b bytes.Buffer
	html.Render(&b, n)
	s := b.String()
	return s
}

func nodesToHtml(nodes []*html.Node) string {
	var b bytes.Buffer
	for _, n := range nodes {
		html.Render(&b, n)
	}

	return b.String()
}

func saveToFile(name string, text string) {
	err := os.WriteFile(name, []byte(text), 0644)
	if err != nil {
		panic(err)
	}
}
