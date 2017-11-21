// Tests for challenge

package main

// Import decent testing
// sidenote - who writes a language without assertions!?!?
import (
	"testing"
	"github.com/stretchr/testify/require"
	"os"
)

// FIXME - brittle! Many tests break if the codes change.

func TestInvalidCode(t *testing.T) {
	// Basic check - completely invalid code
	require.False(t,isValid("NOTACODE"))

	// Double-check that the valid code referenced by
	// the tests below actually is valid
	require.True(t,isValid("CODE1"),"Oh dear, looks like CODE1 isn't a valid code anymore")

	// Code that is a superset of a valid code
	require.False(t,isValid("CODE12"))

	// Code that is a subset of a valid code
	require.False(t,isValid("CODE"))
}

func TestUsedCode(t *testing.T) {
	// Move any existing savefile before testing
	os.Rename("progress.dat","progress_old.dat")

	// Once we're done, we'll want to move it back
	defer os.Rename("progress_old.dat","progress.dat")

	// Check we have a valid code
	testCode := "CODE1"
	require.True(t,isValid(testCode))

	// Negative test
	require.False(t,codeUsed(testCode))

	// Store it
	err := storeCode(testCode)
	require.NoError(t,err)

	// Now try the same code again
	require.True(t,codeUsed(testCode))
}

// Test can save more than one code
func TestMultipleCodes(t *testing.T) {
	// Move any existing savefile before testing
	os.Rename("progress.dat","progress_old.dat")

	// Once we're done, we'll want to move it back
	defer os.Rename("progress_old.dat","progress.dat")

	// Run through three times, make sure we can save more than one
	// code (and are therefore appending not overwriting the save
	// state)
	var testCodes = [3]string{"CODE1","CODE2","CODE3"}
	for index,code := range testCodes {
		// Check the code is valid
		require.True(t,isValid(code),code + " no longer a valid code.")

		// Store it
		err := storeCode(code)
		require.NoError(t,err)

		// Confirm code now stored
		require.True(t,codeUsed(code))

		// Test getNumCodes works properly while we're at it
		require.Equal(t,index+1,getNumCodes())
	}

	// Confirm that all codes are still stored
	for _,code := range testCodes {
		require.True(t,codeUsed(code),code + " not stored!")
	}
}

// Test adding random stuff to savefile doesn't show new clues
func TestGetNumCodesJunkInSaveFile(t *testing.T) {
	// Move any existing savefile before testing
	os.Rename("progress.dat","progress_old.dat")

	// Once we're done, we'll want to move it back
	defer os.Rename("progress_old.dat","progress.dat")

	// Quickly check getNumCodes is right when there's no file
	require.Equal(t,0,getNumCodes())

	// Check that our test code is valid (I should really be doing this
	// in a setup method by this point!)
	testCode := "CODE1"
	require.True(t,isValid(testCode),testCode + " no longer a valid code.")

	// Store
	err := storeCode(testCode)
	require.NoError(t,err)

	// Confirm 1 code
	require.Equal(t,1,getNumCodes())

	// Manually open the file and add some more content
	f, err := os.OpenFile("progress.dat", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString("ANOTHER CODE...\nLOADS OF CODES, HONEST.ALSO " + testCode + " AGAIN")

	// Should still only be 1 code
	require.Equal(t,1,getNumCodes())

	// Add another valid one, double-check it's still counted
	testCode = "CODE2"
	require.True(t,isValid(testCode),testCode + " no longer a valid code.")
	err = storeCode(testCode)
	require.NoError(t,err)

	// Now there should be two
	require.Equal(t,2,getNumCodes())
}
