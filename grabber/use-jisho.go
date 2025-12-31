package grabber

import (
	"fmt"
	"net/url"
)

func UseJisho(req string) (string, error) {
	u := "https://jisho.org/search/" + url.QueryEscape(req)
	fmt.Println(u)
	doc, err := grab(u, false)
	if err != nil {
		return "", err
	}

	fmt.Println(doc)
	blocks := collectNodes(doc, Tag{"div", "class", "concept_light"})
	// blocks := collectNodes(doc, Tag{"div", "", ""})
	fmt.Println(blocks)
	for _, block := range blocks {
		removeTags(block, []Tag{
			{"div", "class", "concept_light-status"},
			{"a", "class", "light-details_link"},
		})
	}

	// fmt.Println(nodesToHtml(blocks))
	// saveToFile("parsed-j3.html", nodesToHtml(blocks))
	return nodesToHtml(blocks), nil
	// return nodeToHtml(doc), nil
}
