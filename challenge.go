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
)

const numCodes int = 5
var validCodes = [numCodes]string{"CODE1","CODE2","CODE3","CODE4","CODE5"}
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
	mainMenu(true)
}

// Display main menu
func mainMenu(firstTime bool) {
	// Prep for reading from command line
	reader := bufio.NewReader(os.Stdin)
	
	// Unless first time, show pic
	// TODO

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
	}	
}

// Called from main function when player chooses to enter a code
func enterCode() {
	// Prep for reading from command line
	reader := bufio.NewReader(os.Stdin)
	
	prompt(enterCodeContent)
	
	code, _ := reader.ReadString('\n')
	code = strings.TrimSpace(code)
	
	fmt.Println("Thanks for entering '" + code + "'...")
	if isValid(code) {
		if codeUsed(code) {
			fmt.Println("Oh no, looks like you already had that one!")
		} else {
			// New code
			handleNewCode(code)
		}
	} else {
		fmt.Println("Code invalid :(")
	}
	// Options here - show clues unlocked so far (if any), enter another code, back to main menu
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

// Final technical challenge
func writeSanta() {
}
