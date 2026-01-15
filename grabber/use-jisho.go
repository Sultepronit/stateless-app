package grabber

import (
	"net/url"
)

var allowedClasses = map[string]struct{}{
	"concept_light":                {},
	"concept_light-representation": {},
	"meaning-representation_notes": {},
	"supplemental_info":            {},
	"meaning-tags":                 {},
	"meaning-abstract":             {},
	"furigana":                     {},
	"furigana-justify":             {},
	"kanji":                        {},
	"sentences":                    {},
	"japanese":                     {},
}

func UseJisho(req string) (string, error) {
	u := "https://jisho.org/search/" + url.QueryEscape(req)
	// fmt.Println(u)
	doc, err := grab(u, false)
	if err != nil {
		return "", err
	}

	// fmt.Println(doc)
	blocks := collectNodes(doc, Tag{"div", "class", "concept_light"})
	// blocks := collectNodes(doc, Tag{"div", "", ""})
	// fmt.Println(blocks)

	for _, block := range blocks {
		removeTags(block, []Tag{
			{"div", "class", "concept_light-status"},
			{"a", "class", "light-details_link"},
		})

		cleanClasses(block, allowedClasses)
	}

	// fmt.Println(nodesToHtml(blocks))
	return nodesToHtml(blocks), nil
}
