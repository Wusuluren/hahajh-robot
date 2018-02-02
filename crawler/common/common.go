package common

import (
	"hahajh-robot/util/gquery"
	"strings"
)

func GetChildrenText(node *gquery.HtmlNode) string {
	textArry := make([]string, 0)
	for _, node := range node.Children("") {
		if node.Label == "" {
			textArry = append(textArry, node.Text)
		}
	}
	return strings.Trim(strings.Join(textArry, ""), "\t\n\r ")
}

func IsEmptyNode(node *gquery.HtmlNode) bool {
	return node.Label == "" && node.Text == "" && len(node.Attribute) == 0
}
