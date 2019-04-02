package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
)

// Provides the user interface and root of method execution
func main() {
	// Reader that will read from stdin pipe
	reader := bufio.NewReader(os.Stdin)

	// Create user interface
	welcomeScreen := figure.NewFigure("Quinton's Driver Installer", "", true)
	welcomeScreen.Print()

	fmt.Printf("(1)Install Network\n(2)Dell Command Updat\nOr both(ENTER)?\n")
	decision, err := reader.ReadByte()

	if err != nil {
		networkInstall()
		driverInstall()
	} else if decision == 1 {
		networkInstall()
	} else if decision == 2 {
		driverInstall()
	} else {
		fmt.Println("Invalid argument")
		time.Sleep(time.Second * 5)
	}

	// Eventually want to implement: copy and paste MediaCreationTool to the desktop
}

// findExecutible will search current working
// directory for the desired exe file and return its name as a string
func findExecutable(directory string, keywords ...string) (string, bool) {
	// get all files into one variable
	allfiles := GetAllFiles(directory)

	// check through all files for matching keywords
	for _, item := range allfiles {
		isCompleteMatch, _ := checkSubstrings(item, keywords)

		if isCompleteMatch == true {
			return item, true
		}
	}

	return "Item not found", false
}

// getwd gets current working directory
func getwd() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Print(err)
		return "Nope, couldn't get current working directory"
	}

	return dir
}

// GetAllFiles returns all files within a directory as string in a slice
func GetAllFiles(directory string) []string {

	var listOffiles []string

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		print("OOPS, something went wrong with ReadDir()")
		time.Sleep(time.Second * 5)
	}

	for _, item := range files {
		listOffiles = append(listOffiles, item.Name())
	}

	return listOffiles
}

/*
checkSubstrings goes through to see if your chosen string contains all
substrings. returns true if contains all substrings. false otherwise
*/
func checkSubstrings(str string, subs []string) (bool, int) {

	matches := 0
	isCompleteMatch := true

	for _, sub := range subs {
		if strings.Contains(str, sub) {
			matches++
		} else {
			isCompleteMatch = false
		}
	}

	return isCompleteMatch, matches
}

func networkInstall() {
	// get the name of the executable (SDITool)
	sdiExe, isFound := findExecutable(getwd()+"/Scripts", "SDI", "x64")

	// execute the installer
	if isFound {
		cmd := exec.Command(sdiExe, "-autoinstall", "-nogui", "-showconsole", "-autoclose")

		err := cmd.Run()
		if err != nil {
			fmt.Printf("Failed to execute %s\n", sdiExe)
		}
	} else {
		fmt.Println("Error: Could not find executable")
		time.Sleep(time.Second * 5)
	}

}

// driverInstall silently installs the DCU program
func driverInstall() {
	// get the name of the executable (DCU)
	dciExe, isFound := findExecutable(getwd()+"/Scripts", "DCU")

	// execute the installer
	if isFound {
		cmd := exec.Command(dciExe, "/s")

		err := cmd.Run()

		if err != nil {
			fmt.Printf("Failed to execute %s\n", dciExe)
		}
	} else {
		fmt.Println("Error: Could not find executable")
		time.Sleep(time.Second * 5)
	}
}
