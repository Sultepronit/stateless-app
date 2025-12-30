package grabber

import (
	"fmt"
	"net/url"
	"strings"
)

func UseSlovnyk(req string) (string, error) {
	u := "https://slovnyk.ua/index.php?swrd=" + url.QueryEscape(req)
	fmt.Println(u)
	doc, err := grab(u, true)
	if err != nil {
		return "", err
	}

	blocks := collectNodes(doc, Tag{"div", "class", "toggle-content"})

	var sb strings.Builder
	for _, b := range blocks {
		unwrapTags(b, "a")
		sb.WriteString(nodeToHtml(b))
	}

	// saveToFile("parsed-s1.html", sb.String())
	return sb.String(), nil
}
