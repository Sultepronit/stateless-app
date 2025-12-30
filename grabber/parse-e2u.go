package grabber

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func isHeaderRight(header string, expected string) bool {
	spl := regexp.MustCompile(`[0-9]|\s\(-`)
	header = spl.Split(header, 2)[0]
	repl := strings.NewReplacer("|", "", "́", "")
	header = repl.Replace(header)
	// header = strings.TrimSpace(header)
	return strings.EqualFold(header, expected)
}

func isArticleMain(article *html.Node, query string) bool {
	// b := findNode(article, "b")
	b := findNode(article, Tag{"b", "", ""})
	if b == nil {
		return false
	}

	header := getTextContent(b)

	return isHeaderRight(header, query)
}

func UseE2u(req string) (string, error) {
	u := "https://e2u.org.ua/s?w=" + url.QueryEscape(req) + "&dicts=all&highlight=on&filter_lines=on"
	fmt.Println(u)
	doc, err := grab(u, true)
	if err != nil {
		return "", err
	}

	tds := collectNodes(doc, Tag{"td", "", ""})

	articles := map[string][]*html.Node{
		"main":    {},
		"other":   {},
		"context": {},
	}
	for _, tag := range tds {
		tag.Data = "div"

		if checkAttribute(tag, "class", "result_row") {
			articles["context"] = append(articles["context"], tag)
		} else if isArticleMain(tag, "apple") {
			articles["main"] = append(articles["main"], tag)
		} else {
			articles["other"] = append(articles["other"], tag)
		}
	}

	order := []string{"main", "other", "context"}
	var sb strings.Builder
	for _, group := range order {
		sb.WriteString(`<article class="`)
		sb.WriteString(group)
		sb.WriteString(`">`)
		sb.WriteString(nodesToHtml(articles[group]))
		sb.WriteString(`</article>`)
	}

	return sb.String(), nil
}
