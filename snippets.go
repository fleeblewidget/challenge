// Snippets revealed for correct codes
// NB need to have the same number of snippets as codes

package main

import "fmt"

func showSnippet(snippet int) {
	snippetText := ""
	switch snippet {
	case 1:
		snippetText = `
SNIPPET ONE PLACEHOLDER
`
	case 2:
		snippetText = `
SNIPPET TWO PLACEHOLDER
`
	case 3:
		snippetText = `
SNIPPET THREE PLACEHOLDER
`
	case 4:
		snippetText = `
SNIPPET FOUR PLACEHOLDER
`
	case 5:
		snippetText = `
SNIPPET FIVE PLACEHOLDER
`
	}

	fmt.Println(snippetText)
}
