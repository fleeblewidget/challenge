// Snippets revealed for correct codes
// NB need to have the same number of snippets as codes

package main

import "fmt"

func showSnippet(snippet int) {
	snippetText := ""
	switch snippet {
	case 1:
		snippetText = `
Q: Why does Santa go down the chimney on Christmas Eve?
A: It 'soots' his style!

The first rule of writing to Santa is:

All letters must start by listing the child's name, age and location (on one line)
`
	case 2:
		snippetText = `
Q: What do you call Santa's little helpers?
A: Subordinate clauses!

The second rule of writing to Santa is:

Names must be given in the format "Surname, Firstname, Initials"
`
	case 3:
		snippetText = `
Q: Where does a snowman put his birthday candles?
A: On his birthday flake!

The third rule of writing to Santa is:

Age must be given in months (e.g. 57 months)
`
	case 4:
		snippetText = `
Q: What did the Atlantic Ocean give the Pacific at Christmas?
A: Nothing, they just waved.

The fourth rule of writing to Santa is:

Location must be given in decimal degrees latitude and longitude
`
	case 5:
		snippetText = `
Q: What kind of music do elves like best?
A: "Wrap" music!

The fifth rule of writing to Santa is:

Items must be listed in ascending order of volume (to help with packing)
`

	case 6:
		snippetText = `
Q: How does Christmas Day end?
A: With the letter Y!

The sixth rule of writing to Santa is:

Each item must be requested on a separate line, in the format "Please may I have a(n)..."
`
	}

	fmt.Println(snippetText)
}
