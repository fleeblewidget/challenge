// Technical challenge platform
//
// NB writing this was used as an excuse to learn some Go - probably not great
// code!
//
// Author: fleeblewidget
// Available freely for use, modification and distribution under the GNU
// General Public Licence
package main

import (
	"fmt"
	"io/ioutil"
	"bufio"
	"os"
	"strings"
	"regexp"
	"time"
)

const numCodes int = 6
var validCodes = [numCodes]string{"32139","WDYMU","eCeyo","14849","STWIH","SNGED"}
const saveFile string = "progress.dat"

// Print a specified message and display themed prompt
func prompt(message string) {
	fmt.Println(message)
	fmt.Print("ELF>")
}

// Check whether a code is valid
func isValid(code string) bool {
	for _,validCode := range validCodes {
		if validCode == code {
			return true
		}
	}
	return false
}

// Check whether a code has been used
// Use after isValid
func codeUsed(code string) bool {
	dat, err := ioutil.ReadFile(saveFile)
	if err != nil {
		// No file, so code hasn't been used
		return false
	}
	if strings.Contains(string(dat),code) {
		return true
	} else {
		return false
	}
}

// Store a correctly-entered code in savedata
func storeCode(code string) error {
	f, err := os.OpenFile(saveFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(code)
	return err
}

// Find number of valid codes entered so far
func getNumCodes() int {
	dat, err := ioutil.ReadFile(saveFile)
	savedata := string(dat)
	if err != nil {
		// No file, no codes
		return 0
	}
	numHits := 0
	for _, validCode := range validCodes {
		if strings.Contains(savedata,validCode) {
			numHits++
		}
	}
	return numHits
}

func main() {
	// Display welcome message
	fmt.Println(header)

	// Show main menu
	mainMenu()
}

// Display main menu
func mainMenu() {
	// Prep for reading from command line
	reader := bufio.NewReader(os.Stdin)
	firstTime := true

	for {
		// Unless first time, show pic
		if (!firstTime) {
			fmt.Print(mainMenuPic)
		}

		// Text
		prompt(mainMenuContent)
	
		// Get option
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		// Run subroutine for chosen option
		if (command == "e" || command == "E") {
			enterCode()
		}
		if (command == "w" || command == "W") {
			writeSanta()
			return
		}
		if (command == "q" || command == "Q") {
			break
		}
		firstTime = false
	}

	// End of program
	fmt.Println(goodbyeContent)
}

// Called from main function when player chooses to enter a code
func enterCode() {
	// Prep for reading from command line
	reader := bufio.NewReader(os.Stdin)
	
	// Non-terminating loop - keep entering codes until the user chooses to go back to the main menu
	for {
		prompt(enterCodeContent)
	
		code, _ := reader.ReadString('\n')
		code = strings.TrimSpace(code)
	
		fmt.Println("Thanks for entering '" + code + "'...")
		if isValid(code) {
			if codeUsed(code) {
				fmt.Println("Oh no! You tried that code already! That must be as disappointing as getting two identical jokes out of the same set of crackers! :-( Try again!")
			} else {
				// New code
				handleNewCode(code)
			}
		} else {
			fmt.Println("Code invalid :(")
		}
		
		// Options here - show clues unlocked so far (if any), enter another code, back to main menu
		fmt.Println("\nWhat next?\n")
		numCodes := getNumCodes()
		if (numCodes > 0) {
			fmt.Printf("S)ee snippets unlocked so far - %d codes\n", numCodes)
		}
		prompt("E)nter another code\nR)eturn to main menu")
		
		// Get option
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		// Handle option
		if (command == "s" || command == "S") {
			// Show all unlocked snippets (then go back to main menu)
			showSnippets(numCodes)
			return
		}
		if (command == "r" || command == "R") {
			// Break from loop and return to main menu
			return
		}
		if (command == "e" || command == "E") {
			// Go round loop again, try another code
		}
	}
}

// Handle newly-entered code
func handleNewCode(code string) {
	// Show message
	fmt.Println(validCodeContent)

	// Add this code to the list stored in the file - exit on
	// error, don't show a snippet if can't save
	err := storeCode(code)
	if err != nil {
		fmt.Println("Uh-oh - error saving data.")
		fmt.Println(err)
		os.Exit(1)
	}

	// Show correct snippet
	showSnippet(getNumCodes())
}

// Show all clues unlocked so far
func showSnippets(numCodes int) {
	i := 1
	for i <= numCodes {
		fmt.Printf("\nSnippet %d\n\n",i)
		showSnippet(i)
		i++
	}
}

// Final technical challenge
func writeSanta() {
	// Prep for reading from command line
	scanner := bufio.NewScanner(os.Stdin)
	
	// Header, request re-written letter
	prompt(writeSantaContent)

	// Get letter - need to keep reading until we get the
	// end marker
	var letter []string
	var line string
	for {
		for scanner.Scan() {
			line = scanner.Text()

			if strings.TrimSpace(line) == "<END>" {
				break
			}
			letter = append(letter, line)
		}
		if strings.TrimSpace(line) == "<END>" {
			break
		}
	}

	// Check whether letter passes muster
	if (len(letter) > 0) {
		pass := checkLetter(letter)
		if (pass) {
			fmt.Printf(successContent)
			time.Sleep(10000 * time.Millisecond)
		} else {
			fmt.Printf("\nSorry, that's not right.\n")
		}
	}
}

// Check format of letter
func checkLetter(letter []string) bool {
	// Check header
	match, _ := regexp.MatchString("(?i)scott,? ?kevin,? ?m,? ?81,? ?.?51.768[0-9]*, ?-1.227[0-9]*.?", letter[0])
	if (!match) {
		fmt.Printf("Error: Name/Age/Location not accepted")
		return false
	}

	// Check number of lines
	if (len(letter) != 6) {
		fmt.Printf("Error: wrong length")
		return false
	}
	
	// Check order
	match, _ = regexp.MatchString("(?i).*paw patrol.*",letter[1])
	if (match) {
		match, _ = regexp.MatchString("(?i).*sea ?monkeys.*",letter[2])
	}
	if (match) {
		match, _ = regexp.MatchString("(?i).*colour.*pencil.*",letter[3])
	}
	if (match) {
		match, _ = regexp.MatchString("(?i).*bike.*",letter[4])
	}
	if (match) {
		match, _ = regexp.MatchString("(?i).*elephant.*",letter[5])
	}

	if (!match) {
		fmt.Printf("Error: order not accepted")
		return false
	}

	// Check for "please may I"
	i := 1
	for i <= 5 {
		match, _ = regexp.MatchString("(?i)please may i have.*",letter[i])
		if (!match) {
			fmt.Printf("Error: manners!")
			return false
		}
		i++
		
	}
	
	return true
}
