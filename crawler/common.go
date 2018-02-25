package crawler

import (
	"github.com/wusuluren/gquery"
	"strings"
)

func getChildrenText(node *gquery.HtmlNode) string {
	textArry := make([]string, 0)
	for _, node := range node.Children("") {
		if node.Label == "" {
			textArry = append(textArry, node.Text)
		}
	}
	return strings.Trim(strings.Join(textArry, ""), "\t\n\r ")
}

func isEmptyNode(node *gquery.HtmlNode) bool {
	return node.Label == "" && node.Text == "" && len(node.Attribute) == 0
}
