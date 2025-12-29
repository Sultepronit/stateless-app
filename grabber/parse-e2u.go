package grabber

import (
	"fmt"
	"net/url"
)

// func isHeaderRight(header string, expected string) bool {
// 	spl := regexp.MustCompile(`[0-9]|\s\(-`)
// 	header = spl.Split(header, 2)[0]
// 	repl := strings.NewReplacer("|", "", "́", "")
// 	header = repl.Replace(header)
// 	// header = strings.TrimSpace(header)
// 	// fmt.Println("[", header, "]")
// 	// fmt.Println("[", expected, "]")
// 	return strings.EqualFold(header, expected)
// }

// func isArticleMain(article *html.Node, query string) bool {
// 	b := findNode(article, "b")
// 	if b == nil {
// 		return false
// 	}

// 	header := getTextContent(b)
// 	// fmt.Println(header)

// 	return isHeaderRight(header, query)
// }

func UseE2u(req string) {
	fmt.Println(req)
	esc := url.QueryEscape(req)
	fmt.Println(esc)
	url := "https://e2u.org.ua/s?w=" + req + "&dicts=all&highlight=on&filter_lines=on"
	fmt.Println(url)
	// doc, err := grab(url, true)
	// if err != nil {
	// 	panic(err)
	// }

	// tds := make([]*html.Node, 0, 5)
	// tds = collectTheNodes(doc, tds, "td")

	// articles := map[string][]*html.Node{
	// 	"main":    {},
	// 	"other":   {},
	// 	"context": {},
	// }
	// for _, tag := range tds {
	// 	tag.Data = "div"

	// 	if checkAttribute(tag, "class", "result_row") {
	// 		articles["context"] = append(articles["context"], tag)
	// 	} else if isArticleMain(tag, "apple") {
	// 		articles["main"] = append(articles["main"], tag)
	// 	} else {
	// 		articles["other"] = append(articles["other"], tag)
	// 	}
	// }

	// order := []string{"main", "other", "context"}
	// var sb strings.Builder
	// for _, group := range order {
	// 	sb.WriteString(`<article class="`)
	// 	sb.WriteString(group)
	// 	sb.WriteString(`">`)
	// 	sb.WriteString(nodesToHtml(articles[group]))
	// 	sb.WriteString(`</article>`)
	// }

	// saveToFile("parsed3.html", sb.String())
}
